package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("tr > td").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		fmt.Printf("%T", text)
		fmt.Println(text)
	})
}
