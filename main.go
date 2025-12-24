package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"

	"google.golang.org/protobuf/encoding/protowire"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"golang.org/x/sys/unix"

	klauspost_gzip "github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
)

const zeroRatio = 2

type compressor = func([]byte) []byte

func main() {
	compressors := map[string]compressor{
		"noop": func(data []byte) []byte { return data },
		"gzip_std_best_compression": func(data []byte) []byte {
			var b bytes.Buffer
			w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
			if err != nil {
				log.Fatalf("Creating gzip writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		},
		"gzip_klauspost_compression": func(data []byte) []byte {
			var b bytes.Buffer
			w, err := klauspost_gzip.NewWriterLevel(&b, klauspost_gzip.BestCompression)
			if err != nil {
				log.Fatalf("Creating gzip writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		},
		"zstd_best_compression": func(data []byte) []byte {
			var b bytes.Buffer
			w, err := zstd.NewWriter(&b, zstd.WithEncoderCRC(false), zstd.WithEncoderLevel(zstd.SpeedBestCompression))
			if err != nil {
				log.Fatalf("Creating zstd writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		},
	}

	for name, comp := range compressors {
		if err := testAndWrite(name, comp); err != nil {
			log.Printf("Error testing and writing %s: %v", name, err)
		}
	}
}

func testAndWrite(name string, comp compressor) error {
	results := test(comp)

	cdf := results
	var sum uint64
	for i, count := range results {
		sum += count
		cdf[i] = sum
	}

	// print graph of the cdf
	pts := make(plotter.XYs, len(cdf))
	for i, v := range cdf {
		pts[i].X = float64(i) / float64(len(cdf)) * zeroRatio
		pts[i].Y = float64(v) / float64(sum)
	}

	p := plot.New()
	p.Title.Text = name + " CDF"
	p.X.Label.Text = "Reciprocal Compression Ratio"
	p.Y.Label.Text = "CDF"
	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "0"},
		{Value: 0.25, Label: "0.25"},
		{Value: 0.5, Label: "0.5"},
		{Value: 0.75, Label: "0.75"},
		{Value: 1, Label: "1"},
		{Value: 1.25, Label: "1.25"},
		{Value: 1.5, Label: "1.5"},
		{Value: 1.75, Label: "1.75"},
		{Value: 2, Label: "2"},
	})

	line, err := plotter.NewLine(pts)
	if err != nil {
		return fmt.Errorf("creating line plot: %w", err)
	}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, name+"_cdf.png"); err != nil {
		return fmt.Errorf("saving plot: %w", err)
	}
	return nil
}

func countRows(tableName string, db *sql.DB) uint {
	var count uint
	err := db.QueryRow("SELECT COUNT(*) FROM " + tableName).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

const datasetName = "packets.bin"

func test(comp compressor) (buckets [1024]uint64) {
retry:
	{
		datasetFile, err := os.Open(datasetName)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Printf("Dataset %s not found, generating...", datasetName)
				err = generateDatasetBin(datasetName)
				if err != nil {
					log.Fatalf("Generating dataset: %v", err)
				}
				goto retry
			}
			log.Fatalf("Opening dataset: %v", err)
		}
		defer datasetFile.Close()
		dataset := bufio.NewReader(datasetFile)

		start := time.Now()
		var totalRows uint64
		err = binary.Read(dataset, binary.LittleEndian, &totalRows)
		if err != nil {
			log.Fatalf("Reading total rows from dataset: %v", err)
		}

		tasks := make(chan []byte)
		var wg sync.WaitGroup
		cores := runtime.GOMAXPROCS(0)
		wg.Add(cores)
		for range cores {
			go func() {
				defer wg.Done()
				for payload := range tasks {
					compressed := comp(payload)

					compressionRatio := float64(len(compressed)) / float64(len(payload))
					bucket := int(compressionRatio * float64(len(buckets)) / zeroRatio)
					if bucket < 0 {
						bucket = 0
					}
					if bucket >= len(buckets) {
						bucket = len(buckets) - 1
					}
					atomic.AddUint64(&buckets[bucket], 1)
				}
			}()
		}
		defer wg.Wait()
		defer close(tasks)

		var done uint64
		totalRows = 50000
		for range totalRows {
			done++
			if done%1000 == 0 {
				elapsed := time.Since(start)
				remaining := time.Duration(float64(elapsed) / float64(done) * float64(totalRows-done))
				log.Printf("Processed %d/%d rows (%.2f%%), elapsed: %s, remaining: %s",
					done, totalRows, float64(done)/float64(totalRows)*100,
					elapsed.Truncate(time.Second), remaining.Truncate(time.Second))
			}

			length, err := dataset.ReadByte()
			if err != nil {
				log.Fatalf("Reading length: %s", err)
			}

			payload := make([]byte, length)
			_, err = io.ReadFull(dataset, payload)
			if err != nil {
				log.Fatalf("Reading payload: %s", err)
			}

			tasks <- payload
		}
		return
	}
}

func generateDatasetBin(filename string) error {
	db, err := sql.Open("sqlite", "file:packets_recovered.db?mode=ro")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.OpenFile(".", unix.O_TMPFILE|os.O_WRONLY, 0622)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	rows, err := db.Query(`SELECT payload FROM packet`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	start := time.Now()
	totalRows := countRows("packet", db)

	w := bufio.NewWriter(file)

	// Number of entries placeholder
	_, err = w.WriteString("\x00\x00\x00\x00\x00\x00\x00\x00")
	if err != nil {
		return fmt.Errorf("writing placeholder to file: %w", err)
	}

	var totalEntries uint64
	var done uint
	for rows.Next() {
		done++
		if done%50000 == 0 {
			elapsed := time.Since(start)
			remaining := time.Duration(float64(elapsed) / float64(done) * float64(totalRows-done))
			log.Printf("Creating binary file %d/%d rows (%.2f%%), elapsed: %s, remaining: %s",
				done, totalRows, float64(done)/float64(totalRows)*100,
				elapsed.Truncate(time.Second), remaining.Truncate(time.Second))
		}

		var payload []byte
		err := rows.Scan(&payload)
		if err != nil {
			log.Fatal(err)
		}

		data, ok := extractLoraPayloadFromMessage(payload)
		if !ok || len(data) == 0 || len(data) > 255 {
			continue
		}

		err = w.WriteByte(byte(len(data)))
		if err != nil {
			return fmt.Errorf("writing entry length to file: %w", err)
		}

		_, err = w.Write(data)
		if err != nil {
			return fmt.Errorf("writing entry to file: %w", err)
		}
		totalEntries++
	}
	err = w.Flush()
	if err != nil {
		return fmt.Errorf("flushing to file: %w", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seeking file: %w", err)
	}

	err = binary.Write(file, binary.LittleEndian, totalEntries)
	if err != nil {
		return fmt.Errorf("writing total entries to file: %w", err)
	}

	syscallConn, err := file.SyscallConn()
	if err != nil {
		return fmt.Errorf("getting syscall connection: %w", err)
	}
	var errr error
	err = syscallConn.Control(func(fd uintptr) {
		errr = unix.Fdatasync(int(fd))
		if errr != nil {
			errr = fmt.Errorf("syncing file: %w", errr)
			return
		}
		errr = unix.Linkat(int(fd), "", unix.AT_FDCWD, filename, unix.AT_EMPTY_PATH)
		if errr != nil {
			errr = fmt.Errorf("linking temp file to %s: %w", filename, errr)
			return
		}
	})
	if err != nil {
		return fmt.Errorf("getting control: %w", err)
	}

	return errr
}

func extractLoraPayloadFromMessage(msg []byte) ([]byte, bool) {
	for {
		num, typ, n := protowire.ConsumeTag(msg)
		if n < 0 {
			return nil, false
		}
		msg = msg[n:]

		if num != 4 { // Data field
			continue
		}
		if typ != protowire.BytesType {
			return nil, false
		}

		data, n := protowire.ConsumeBytes(msg)
		if n < 0 {
			return nil, false
		}
		return data, true
	}
}
