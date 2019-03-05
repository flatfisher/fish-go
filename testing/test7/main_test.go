package main

import (
	"flag"
	"testing"
)

var (
	num int
)

func init() {
	flag.IntVar(&num, "num", 1, "")
}

func IsMultiple(n, m int) bool {
	return n%m == 0
}
func TestSampleOdd(t *testing.T) {
	flag.Parse()
	cases := []struct {
		input int
	}{
		{input: 1},
		{input: 3},
		{input: 5},
		{input: 7},
	}
	for _, c := range cases {
		if IsMultiple(c.input, num) {
			t.Log(c.input)
		} else {
			t.Errorf("%d is not multiple of %d", c.input, num)
		}
	}
}
