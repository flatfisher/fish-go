package main

import (
	"fmt"
	"testing"
)

func makeSlice1(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%03d hello", i))
	}
	return r
}

func makeSlice2(n int) []string {
	r := make([]string, n)
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%03d hello", i))
	}
	return r
}

func BenchmarkMakeSlice1(b *testing.B) {
	b.ResetTimer()
	makeSlice1(b.N)
}

func BenchmarkMakeSlice2(b *testing.B) {
	b.ResetTimer()
	makeSlice2(b.N)
}
