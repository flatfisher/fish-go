package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if "Hello World" != getHelloWorld() {
		t.Fatal("failed test")
	}
}
