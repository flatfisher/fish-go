package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if "Hello World" != getHello() {
		t.Fatal("failed test")
	}
}
