package Decode

import (
	"log"
	"testing"
)

func TestDecode_base64Decode(t *testing.T) {
	tests := []struct {
		name string
		d    *Decode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.base64Decode()
		})
	}
}

func TestDecode_removePadding(t *testing.T) {
	tests := []struct {
		name     string
		d        *Decode
		expected string
	}{
		{name: "onePadding", d: &Decode{Input:"Zm9vMTIgYmFyMzQ="}, expected: "Zm9vMTIgYmFyMzQ"},
		{name: "doublePadding", d: &Decode{Input:"Zm9vMTIgYmFyMzQ=="}, expected: "Zm9vMTIgYmFyMzQ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.removePadding()
		})
		if tt.expected != tt.d.Decoded {
			log.Fatalf("\nexpected: %s\nactual:\t  %s", tt.expected, tt.d.Decoded)
		}
	}
}

func TestDecode_decode(t *testing.T) {
	tests := []struct {
		name string
		d    *Decode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.decode()
		})
	}
}

func TestDecode_convertDecimalToBinary(t *testing.T) {
	tests := []struct {
		name string
		d    *Decode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.convertDecimalToBinary()
		})
	}
}

func TestDecode_convertBinaryToString(t *testing.T) {
	tests := []struct {
		name string
		d    *Decode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.convertBinaryToString()
		})
	}
}
