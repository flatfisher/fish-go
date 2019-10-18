package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	buf := bufio.NewWriter(os.Stdout)
	buf.WriteString("hogehoge\n")
	buf.Flush()
	buf.WriteString("hoge\n")
	buf.Flush()

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})

	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加可能")
	request.Write(os.Stdout)
}
