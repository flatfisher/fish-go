package main

import (
	"fmt"
)

// Sample is sample.
type Sample struct {
	X int
	Y int
}

func exchangeSample(s *Sample) {
	x, y := s.Y, s.X
	s.X = x
	s.Y = y
}

func main() {
	s := Sample{X: 1, Y: 2}
	exchangeSample(&s)
	fmt.Println(s.X)
	fmt.Println(s.Y)
}
