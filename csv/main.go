package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	wr := csv.NewWriter(file)
	wr.Write([]string{"test1", "test2", "test3"})
	wr.Flush()
}
