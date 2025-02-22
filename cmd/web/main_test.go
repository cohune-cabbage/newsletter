// Filename: cmd/web/main_test.go
package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	expected := "Hello, world!"
	got := message()
	if expected != got {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
