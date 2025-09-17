package main

import (
	"fmt"
	"os"
	"strings"
)

type ISO8583Message struct {
	MTI                  string
	PAN                  string
	Amount               string
	TransmissionDateTime string
}

func ValidateFields(fields []string) error {
	if len(fields) != 4 {
		return fmt.Errorf("expected 4 fields, got %d", len(fields))
	}
	return nil
}

func MapFieldsToStruct(fields []string) ISO8583Message {
	return ISO8583Message{
		MTI:                  fields[0],
		PAN:                  fields[1],
		Amount:               fields[2],
		TransmissionDateTime: fields[3],
	}
}

func ParseMessage(msg string) []string {
	fmt.Println(msg)
	parts := strings.Split(msg, "|")
	for i, part := range parts {
		fmt.Printf("Field %d: %s\n", i, part)
	}
	return parts
}

func main() {
	var msg string
	if len(os.Args) < 2 {
		fmt.Println("No message provided - using default")
		msg = "0200|1234567890123456|000000100000|20250916083000"
	} else {
		msg = os.Args[1]
	}
	fields := ParseMessage(msg)
	if err := ValidateFields(fields); err != nil {
		fmt.Println("Validation error:", err)
		return
	}
	parsed := MapFieldsToStruct(fields)
	fmt.Printf("%+v\n", parsed)
}
