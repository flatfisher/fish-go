package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if 100 != calculateTriangleArea(10, 20) {
		t.Fatal("failed test")
	}
}
