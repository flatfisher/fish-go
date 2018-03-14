package main

import (
	"fmt"
)

//Location is location.
type Location struct {
	Lat, Long float64
}

var m map[string]Location

func main() {
	m = make(map[string]Location)
	m["Tokyo"] = Location{
		35.6895, 139.6917,
	}

	fmt.Println(m["Tokyo"])
}
