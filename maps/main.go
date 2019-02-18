package main

import (
	"fmt"
)

//Location is location.
type Location struct {
	Lat, Long float64
}

var locations = map[string]Location{
	"Tokyo":  {35.6895, 139.6917},
	"Miyagi": {38.2688, 40.8721},
}

func main() {
	fmt.Println(locations)
	delete(locations, "Tokyo")
	v, ok := locations["Tokyo"]
	fmt.Println("v?", v)
	fmt.Println("Present?", ok)

	var tmp []string
	tmp = append(tmp, "hello")
	tmp = append(tmp, "hello")
	tmp = append(tmp, "world")
	c := wordCount(tmp)
	fmt.Println(c)
}

func wordCount(a []string) map[string]int {
	c := make(map[string]int)
	for _, word := range a {
		c[word]++
	}
	return c
}
