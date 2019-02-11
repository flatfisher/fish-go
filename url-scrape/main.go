package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	content := []string{}
	s := ""
	for y := 7; y <= 18; y++ {
		year := getString(y)
		for m := 1; m <= 12; m++ {
			month := getString(m)
			url := s + year + "-" + month + ".html"
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			z := html.NewTokenizer(res.Body)
			for z.Token().Data != "html" {
				tt := z.Next()
				if tt == html.StartTagToken {
					t := z.Token()
					if t.Data == "td" {
						inner := z.Next()
						if inner == html.TextToken {
							text, err := sJisttoUtf8((string)(z.Text()))
							if err != nil {
								log.Fatal(err)
							}
							t := strings.TrimSpace(text)
							content = append(content, t)
						}
					}
				}
			}
		}
		break
	}
	file, err := os.Create("sample.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	wr := csv.NewWriter(file)
	wr.Write(content)
	wr.Flush()
}

func getString(n int) string {
	if n >= 10 {
		return fmt.Sprintf("%d", n)
	}
	return fmt.Sprintf("0%d", n)
}

func sJisttoUtf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
