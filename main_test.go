package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	msg := Hello()
	expected := "Hello World!"
	if msg != expected {
		t.Fatalf(`Hello() = %q, want %q`, msg, expected)
	}
}
