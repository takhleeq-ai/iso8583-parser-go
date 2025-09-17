package main

import (
	"testing"
)

func TestParseMessage(t *testing.T) {
	msg := "0200|1234567890123456|000000100000|20250916083000"
	if len(ParseMessage(msg)) != 4 {
		t.Errorf("Expected 4 fields, got %d", len(ParseMessage(msg)))
	}
}
