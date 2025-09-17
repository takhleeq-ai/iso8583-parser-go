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
func TestVatidateFields(t *testing.T) {
	validFields := []string{"0200", "1234567890123456", "000000100000", "20250916083000"}
	invalidFields := []string{"0200", "1234567890123456"}

	if err := ValidateFields(validFields); err != nil {
		t.Errorf("Expected no error for valid fields, got %v", err)
	}
	if err := ValidateFields(invalidFields); err == nil {
		t.Errorf("Expected error for invalid fields, got none")
	}
}

func TestMapFieldsToStruct(t *testing.T) {
	fields := []string{"0200", "1234567890123456", "000000100000", "20250916083000"}
	msg := MapFieldsToStruct(fields)

	if msg.MTI != "0200" || msg.PAN != "1234567890123456" || msg.Amount != "000000100000" || msg.TransmissionDateTime != "20250916083000" {
		t.Errorf("Mapping failed, got %+v", msg)
	}
}
