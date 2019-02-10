package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	doc, err := goquery.NewDocument("")
	if err != nil {
		fmt.Print("url scarapping failed")
	}

	doc.Find("table tr td").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		msg, err := sJisttoUtf8(text)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	})
}

func sJisttoUtf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
