package main

import (
	"encoding/csv"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	f, err := os.Open("url.csv")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		qrCode, _ := qr.Encode(record[1], qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, 200, 200)
		file, _ := os.Create("qrcodes/" + string(record[0]) + ".png")
		defer file.Close()
		png.Encode(file, qrCode)
	}
}
