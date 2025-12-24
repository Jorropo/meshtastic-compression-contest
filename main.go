package main

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"

	"google.golang.org/protobuf/encoding/protowire"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const zeroRatio = 2

type compressor = func([]byte) []byte

func main() {
	db, err := sql.Open("sqlite", "file:packets_recovered.db?mode=ro")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results := test(db, func(data []byte) []byte {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		w.Write(data)
		w.Close()
		return b.Bytes()
	})

	cdf := results
	var sum uint
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
	p.Title.Text = "Compression Ratio CDF"
	p.X.Label.Text = "Reprocial Compression Ratio"
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
		log.Fatal(err)
	}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "compression_cdf.png"); err != nil {
		log.Fatal(err)
	}
}

func countRows(tableName string, db *sql.DB) uint {
	var count uint
	err := db.QueryRow("SELECT COUNT(*) FROM " + tableName).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func test(db *sql.DB, comp compressor) (buckets [1024]uint) {
	start := time.Now()
	totalRows := countRows("packet", db)

	rows, err := db.Query(`SELECT payload FROM packet`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var payload []byte
	var done uint
	for rows.Next() {
		done++
		if done%50000 == 0 {
			elapsed := time.Since(start)
			remaining := time.Duration(float64(elapsed) / float64(done) * float64(totalRows-done))
			log.Printf("Processed %d/%d rows (%.2f%%), elapsed: %s, remaining: %s",
				done, totalRows, float64(done)/float64(totalRows)*100,
				elapsed.Truncate(time.Second), remaining.Truncate(time.Second))
		}

		err := rows.Scan(&payload)
		if err != nil {
			log.Fatal(err)
		}

		data, ok := extractLoraPayloadFromMessage(payload)
		if !ok || len(data) == 0 {
			continue
		}

		compressed := comp(data)

		compressionRatio := float64(len(compressed)) / float64(len(data))
		bucket := int(compressionRatio * float64(len(buckets)) / zeroRatio)
		if bucket < 0 {
			bucket = 0
		}
		if bucket >= len(buckets) {
			bucket = len(buckets) - 1
		}
		buckets[bucket]++
	}
	return
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
