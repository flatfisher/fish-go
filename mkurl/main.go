package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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
			if res.StatusCode == 200 {
				fmt.Println(url)
			}
		}
	}
}

func getString(n int) string {
	if n >= 10 {
		return fmt.Sprintf("%d", n)
	}
	return fmt.Sprintf("0%d", n)
}
