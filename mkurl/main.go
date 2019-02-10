package main

import "fmt"

func main() {
	s := ""
	for y := 7; y <= 18; y++ {
		year := getString(y)
		for m := 1; m <= 12; m++ {
			month := getString(m)
			url := s + year + "-" + month + ".html"
			fmt.Println(url)
		}
	}
}

func getString(n int) string {
	if n >= 10 {
		return fmt.Sprintf("%d", n)
	}
	return fmt.Sprintf("0%d", n)
}
