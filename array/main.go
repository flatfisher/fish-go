package main

import "fmt"

type Port struct {
	Latitude float64
	Logitude float64
}

var PORTS = []Port{
	Port{Latitude: 4.0, Logitude: 4},
	Port{Latitude: 44.0, Logitude: 4},
}

func main() {
	fmt.Println(PORTS[0].Latitude)
}
