package main

import (
	"bufio"
	"bytes"
	"cmp"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"

	"google.golang.org/protobuf/encoding/protowire"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"golang.org/x/sys/unix"

	"github.com/cespare/go-smaz"
	cloudflare_lz4 "github.com/cloudflare/golz4"
	"github.com/inkyblackness/res/compress/rle"
	klauspost_flate "github.com/klauspost/compress/flate"
	klauspost_gzip "github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/s2"
	"github.com/klauspost/compress/snappy"
	klauspost_zlib "github.com/klauspost/compress/zlib"
	"github.com/pierrec/lz4/v4"
	shoco_models "github.com/tmthrgd/shoco/models"

	"github.com/Jorropo/meshtastic-compression-contest/unishox2"
)

const zeroRatio = 2

type compressionInput struct {
	from, to uint32
	data     []byte
}
type compressor = func(compressionInput) []byte

const generateTrainingDataset = false

func main() {
	compressors := map[string]compressor{
		"noop": compressJustBytes(func(data []byte) []byte { return data }),
		"gzip_std": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
			if err != nil {
				log.Fatalf("Creating gzip writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"gzip_klauspost": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := klauspost_gzip.NewWriterLevel(&b, klauspost_gzip.BestCompression)
			if err != nil {
				log.Fatalf("Creating gzip writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"flate_std": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := flate.NewWriter(&b, flate.BestCompression)
			if err != nil {
				log.Fatalf("Creating flate writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"flate_klauspost": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := klauspost_flate.NewWriter(&b, klauspost_flate.BestCompression)
			if err != nil {
				log.Fatalf("Creating flate writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"zlib_std": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
			if err != nil {
				log.Fatalf("Creating zlib writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"zlib_klauspost": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w, err := klauspost_zlib.NewWriterLevel(&b, klauspost_zlib.BestCompression)
			if err != nil {
				log.Fatalf("Creating zlib writer: %v", err)
			}
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"lzw_std": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w := lzw.NewWriter(&b, lzw.LSB, 8)
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"s2_klauspost": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w := s2.NewWriter(&b, s2.WriterBestCompression(), s2.WriterConcurrency(1))
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"snappy_klauspost": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w := snappy.NewBufferedWriter(&b)
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"unishox2_meshtastic": compressorOnlyTextMessageAppContent(func(data []byte) []byte {
			// copied from https://github.com/meshtastic/firmware/blob/3a7093a973c1b16d2d978576f1f880ed4c8d7386/src/mesh/Router.cpp#L570
			compressedPayload, err := unishox2.CompressSimple(data)
			if err != nil {
				log.Fatalf("Compressing with unishox2: %v", err)
			}
			return compressedPayload
		}),
		"rle_inkyblackness": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			rle.Compress(&b, data)
			return b.Bytes()
		}),
		"lz4_pierrec": compressJustBytes(func(data []byte) []byte {
			var b bytes.Buffer
			w := lz4.NewWriter(&b)
			w.Apply(lz4.BlockChecksumOption(false), lz4.ChecksumOption(false), lz4.CompressionLevelOption(lz4.Level9), lz4.ConcurrencyOption(1))
			w.Write(data)
			w.Close()
			return b.Bytes()
		}),
		"lz4_cloudflare": compressJustBytes(func(data []byte) []byte {
			r := make([]byte, cloudflare_lz4.CompressBound(data))
			n, err := cloudflare_lz4.Compress(data, r)
			if err != nil {
				panic(err)
			}
			return r[:n]
		}),
		"lz4_cloudflareHC": compressJustBytes(func(data []byte) []byte {
			r := make([]byte, cloudflare_lz4.CompressBound(data))
			n, err := cloudflare_lz4.CompressHCLevel(data, r, 16)
			if err != nil {
				panic(err)
			}
			return r[:n]
		}),
		"smaz_cespare":         compressJustBytes(smaz.Compress),
		"smaz_cespare_Jorropo": compressorOnlyTextMessageAppContent(smaz.Compress),
		"shoco_WordsEn_tmthrgd": compressJustBytes(func(data []byte) []byte {
			return shoco_models.WordsEn().ProposedCompress(data)
		}),
		"shoco_WordsEn_tmthrgd_Jorropo": compressorOnlyTextMessageAppContent(func(data []byte) []byte {
			return shoco_models.WordsEn().ProposedCompress(data)
		}),
		"shoco_TextEn_tmthrgd": compressJustBytes(func(data []byte) []byte {
			return shoco_models.TextEn().ProposedCompress(data)
		}),
		"shoco_TextEn_tmthrgd_Jorropo": compressorOnlyTextMessageAppContent(func(data []byte) []byte {
			return shoco_models.TextEn().ProposedCompress(data)
		}),
		"shoco_FilePath_tmthrgd": compressJustBytes(func(data []byte) []byte {
			return shoco_models.FilePath().ProposedCompress(data)
		}),
		"shoco_FilePath_tmthrgd_Jorropo": compressorOnlyTextMessageAppContent(func(data []byte) []byte {
			return shoco_models.FilePath().ProposedCompress(data)
		}),
		"shoco_Emails_tmthrgd": compressJustBytes(func(data []byte) []byte {
			return shoco_models.Emails().ProposedCompress(data)
		}),
		"shoco_Emails_tmthrgd_Jorropo": compressorOnlyTextMessageAppContent(func(data []byte) []byte {
			return shoco_models.Emails().ProposedCompress(data)
		}),
		"snowflake_Jorropo": explodePacketForPortnumPayloadSubstitution(compressPerPortnumTuned),
	}

	type resultPair struct {
		name                       string
		avg, avgOnlyTextMessageApp float64
	}

	var results = []resultPair{}

	const nameOnlyTextMessageAppSuffix = " only TEXT_MESSAGE_APP"

	for name, comp := range compressors {
		avg, err := testAndWrite(name, comp, false)
		if err != nil {
			log.Fatalf("Error testing and writing %s: %v", name, err)
		}

		nameOnlyTextMessageApp := name + nameOnlyTextMessageAppSuffix

		avgOnlyTextMessageApp, err := testAndWrite(nameOnlyTextMessageApp, comp, true)
		if err != nil {
			log.Fatalf("Error testing and writing %s: %v", nameOnlyTextMessageApp, err)
		}
		results = append(results, resultPair{name: name, avg: avg, avgOnlyTextMessageApp: avgOnlyTextMessageApp})
	}

	slices.SortFunc(results, func(a, b resultPair) int {
		r := cmp.Compare(a.avgOnlyTextMessageApp, b.avgOnlyTextMessageApp)
		if r == 0 {
			r = cmp.Compare(a.name, b.name)
		}
		return r
	})

	var README bytes.Buffer

	README.WriteString(`# Meshtastic Compression Showdown

This project contain benchmarks of various compression algorithms applied on a data of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
One **bellow** 1 means the compressed data is **smaller** than the uncompressed data.

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------------------------------|
`)

	for _, r := range results {
		fmt.Fprintf(&README, "| `%s` | %.4f |\n", r.name, r.avgOnlyTextMessageApp)
	}

	README.WriteString(`
| Compressor | Average Reciprocal Compression Ratio |
|------------|--------------------------------------|
`)

	slices.SortFunc(results, func(a, b resultPair) int {
		r := cmp.Compare(a.avg, b.avg)
		if r == 0 {
			r = cmp.Compare(a.name, b.name)
		}
		return r
	})

	for _, r := range results {
		fmt.Fprintf(&README, "| `%s` | %.4f |\n", r.name, r.avg)
	}

	README.WriteString(`
## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

`)

	for _, r := range results {
		fmt.Fprintf(&README, "### `%s`\n\n![%s CDF](graphs/%s_cdf.png)\n\n![%s CDF](graphs/%s_cdf.png)\n\n", r.name, r.name+nameOnlyTextMessageAppSuffix, strings.ReplaceAll(r.name+nameOnlyTextMessageAppSuffix, " ", "_"), r.name, strings.ReplaceAll(r.name, " ", "_"))
	}

	f, err := os.OpenFile("README.md", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening README.md: %v", err)
	}
	defer f.Close()

	if _, err := README.WriteTo(f); err != nil {
		log.Fatalf("Error writing to README.md: %v", err)
	}
}

func compressJustBytes(comp func([]byte) []byte) compressor {
	return func(input compressionInput) []byte {
		return comp(input.data)
	}
}

func compressorOnlyTextMessageAppContent(comp func([]byte) []byte) compressor {
	return explodePacketForPortnumPayloadSubstitution(func(portnum uint64, from, to uint32, payload []byte) (newPortnum uint64, newPayload []byte, changed bool) {
		if portnum != TEXT_MESSAGE_APP {
			return 0, nil, false
		}
		newPayload = comp(payload)
		return COMPRESSED, newPayload, true
	})
}

func testAndWrite(name string, comp compressor, onlyTextMessageApp bool) (avg float64, err error) {
	log.Printf("Testing compressor: %s", name)
	results, avg := test(comp, onlyTextMessageApp)

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
		return 0, fmt.Errorf("creating line plot: %w", err)
	}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "graphs/"+strings.ReplaceAll(name, " ", "_")+"_cdf.png"); err != nil {
		return 0, fmt.Errorf("saving plot: %w", err)
	}
	return avg, nil
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
const tryLimit = 10000

var cachedDataset []byte
var loadDatasetOnce sync.Once

func test(comp compressor, onlyTextMessageApp bool) (buckets [1024]uint64, avg float64) {
	loadDatasetOnce.Do(func() {
	retry:
		{
			var err error
			cachedDataset, err = os.ReadFile(datasetName)
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
		}
	})
	dataset := cachedDataset

	start := time.Now()
	totalRows := binary.LittleEndian.Uint64(dataset[:8])
	dataset = dataset[8:]

	var sumCompressed uint64
	tasks := make(chan compressionInput)
	var wg sync.WaitGroup
	cores := runtime.GOMAXPROCS(0)
	wg.Add(cores)
	for range cores {
		go func() {
			defer wg.Done()
			for payload := range tasks {
				compressed := comp(payload)

				compressionRatio := float64(len(compressed)) / float64(len(payload.data))
				bucket := int(compressionRatio * float64(len(buckets)) / zeroRatio)
				if bucket < 0 {
					bucket = 0
				}
				if bucket >= len(buckets) {
					bucket = len(buckets) - 1
				}
				atomic.AddUint64(&buckets[bucket], 1)
				atomic.AddUint64(&sumCompressed, uint64(len(compressed)))
			}
		}()
	}

	var done, sumUncompressed uint64
	for range totalRows {
		from, to := binary.LittleEndian.Uint32(dataset[:4]), binary.LittleEndian.Uint32(dataset[4:8])
		dataset = dataset[8:]

		length := dataset[0]
		dataset = dataset[1:]

		payload := dataset[:length]
		dataset = dataset[length:]

		if onlyTextMessageApp {
			portnum, _, _, _, ok := extractPortnumAndPayloadFromDecoded(payload)
			if !ok {
				panic("unreachable")
			}
			if portnum != TEXT_MESSAGE_APP {
				continue
			}
		}

		tasks <- compressionInput{from: from, to: to, data: payload}
		sumUncompressed += uint64(len(payload))

		done++
		if done%2000 == 0 {
			elapsed := time.Since(start)
			remaining := time.Duration(float64(elapsed) / float64(done) * float64(totalRows-done))
			log.Printf("Processed %d/%d rows (%.2f%%), elapsed: %s, remaining: %s",
				done, totalRows, float64(done)/float64(totalRows)*100,
				elapsed.Truncate(time.Second), remaining.Truncate(time.Second))
		}
	}
	close(tasks)
	wg.Wait()
	avg = float64(sumCompressed) / float64(sumUncompressed)
	return
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

	var seed [32]byte
	copy(seed[:], "Jorropo")
	rng := rand.New(rand.NewChaCha8(seed))

	var totalEntries uint64
	var done, trainPackets, trainTextMessage uint
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

		from, to, data, ok := extractLoraPayloadFromMessage(payload)
		if !ok || len(data) == 0 || len(data) > 255 {
			continue
		}

		roll := rng.UintN(totalRows)
		// randomly pick messages for training or benchmark set
		if roll < tryLimit {
			binary.Write(w, binary.LittleEndian, from)
			binary.Write(w, binary.LittleEndian, to)

			err = w.WriteByte(byte(len(data)))
			if err != nil {
				return fmt.Errorf("writing entry length to file: %w", err)
			}

			_, err = w.Write(data)
			if err != nil {
				return fmt.Errorf("writing entry to file: %w", err)
			}
			totalEntries++
		} else if generateTrainingDataset {
		storeTrainingPacket:
			{
				f, err := os.OpenFile("train/packets/"+strconv.FormatUint(uint64(trainPackets), 36), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						err = os.MkdirAll("train/packets", 0755)
						if err != nil {
							return fmt.Errorf("creating train/packets directory: %w", err)
						}
						goto storeTrainingPacket
					}
					return fmt.Errorf("creating individual packet file: %w", err)
				}
				_, err = f.Write(data)
				if err != nil {
					return fmt.Errorf("writing individual packet file: %w", err)
				}
				err = f.Close()
				if err != nil {
					return fmt.Errorf("closing individual packet file: %w", err)
				}
				trainPackets++
			}

		storePortnumAndPayloadTraining:
			{
				portnum, _, payload, _, ok := extractPortnumAndPayloadFromDecoded(data)
				if ok {
					f, err := os.OpenFile("train/"+strconv.FormatUint(uint64(portnum), 10)+"/"+strconv.FormatUint(uint64(trainTextMessage), 36), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0622)
					if err != nil {
						if errors.Is(err, os.ErrNotExist) {
							err = os.MkdirAll("train/"+strconv.FormatUint(uint64(portnum), 10), 0755)
							if err != nil {
								return fmt.Errorf("creating train/packets directory: %w", err)
							}
							goto storePortnumAndPayloadTraining
						}
						return fmt.Errorf("creating individual text message file: %w", err)
					}
					_, err = f.Write(payload)
					if err != nil {
						return fmt.Errorf("writing individual text message file: %w", err)
					}
					err = f.Close()
					if err != nil {
						return fmt.Errorf("closing individual text message file: %w", err)
					}
				}
				trainTextMessage++
			}
		}
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

func extractLoraPayloadFromMessage(msg []byte) (from, to uint32, payload []byte, ok bool) {
	for len(msg) > 0 {
		num, typ, n := protowire.ConsumeTag(msg)
		if n < 0 {
			return 0, 0, nil, false
		}
		msg = msg[n:]

		switch num {
		case 1: // From field
			if typ != protowire.Fixed32Type {
				return 0, 0, nil, false
			}

			from, n = protowire.ConsumeFixed32(msg)
			if n < 0 {
				return 0, 0, nil, false
			}
			msg = msg[n:]
		case 2: // To field
			if typ != protowire.Fixed32Type {
				return 0, 0, nil, false
			}

			to, n = protowire.ConsumeFixed32(msg)
			if n < 0 {
				return 0, 0, nil, false
			}
			msg = msg[n:]
		case 4: // Data field
			if typ != protowire.BytesType {
				return 0, 0, nil, false
			}

			payload, n = protowire.ConsumeBytes(msg)
			if n < 0 {
				return 0, 0, nil, false
			}
			msg = msg[n:]
		default:
			n = protowire.ConsumeFieldValue(num, typ, msg)
			if n < 0 {
				return 0, 0, nil, false
			}
			msg = msg[n:]
		}

	}

	_, _, _, _, ok = extractPortnumAndPayloadFromDecoded(payload)
	if !ok {
		return 0, 0, nil, false
	}

	return from, to, payload, true
}

func extractPortnumAndPayloadFromDecoded(data []byte) (portnum uint64, before, payload, after []byte, ok bool) {
	msg := data
	for len(msg) > 0 {
		num, typ, n := protowire.ConsumeTag(msg)
		if n < 0 {
			return 0, nil, nil, nil, false
		}
		msgBeforeConsumeTag := msg
		msg = msg[n:]
		switch num {
		case 1: // portnum
			if typ != protowire.VarintType {
				return 0, nil, nil, nil, false
			}
			portnum, n = protowire.ConsumeVarint(msg)
			if n < 0 {
				return 0, nil, nil, nil, false
			}
			msg = msg[n:]
			if payload != nil {
				return portnum, before, payload, after, true
			}
		case 2: // payload
			if typ != protowire.BytesType {
				return 0, nil, nil, nil, false
			}
			before = data[:len(data)-len(msgBeforeConsumeTag)]
			payload, n = protowire.ConsumeBytes(msg)
			if n < 0 {
				return 0, nil, nil, nil, false
			}
			after = msg[n:]
			if portnum != 0 {
				return portnum, before, payload, after, true
			}
		default:
			n = protowire.ConsumeFieldValue(num, typ, msg)
			if n < 0 {
				return 0, nil, nil, nil, false
			}
			msg = msg[n:]
		}
	}
	return 0, nil, nil, nil, false
}
