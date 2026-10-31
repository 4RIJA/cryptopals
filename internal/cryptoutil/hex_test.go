package cryptoutil

import (
	"bytes"
	"testing"
)

func TestHexValue(t *testing.T) {
	tests := []struct {
		in   byte
		want byte
	}{
		{'0', 0},
		{'9', 9},
		{'a', 10},
		{'f', 15},
		{'A', 10},
		{'F', 15},
	}
	for _, tt := range tests {
		got, err := hexValue(tt.in)
		if err != nil {
			t.Fatalf("hexValue(%q) devolvio error: %v", tt.in, err)
		}
		if got != tt.want {
			t.Errorf("hexValue(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestHexDecode(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []byte
	}{
		{"vacio", "", []byte{}},
		{"cero", "00", []byte{0x00}},
		{"un byte", "4f", []byte{0x4f}},
		{"mayusculas", "4F", []byte{0x4f}},
		{"varios bytes", "deadbeef", []byte{0xde, 0xad, 0xbe, 0xef}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexDecode(tt.in)
			if err != nil {
				t.Fatalf("hexDecode(%q) devolvio error: %v", tt.in, err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("hexDecode(%q) = %x, want %x", tt.in, got, tt.want)
			}
		})
	}
}
func TestValueHex(t *testing.T) {
	tests := []struct {
		in   byte
		want byte
	}{
		{0, '0'},
		{9, '9'},
		{10, 'a'},
		{15, 'f'},
	}
	for _, tt := range tests {
		got, _ := valueHex(tt.in)
		if got != tt.want {
			t.Errorf("valueHex(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestHexEncode(t *testing.T) {
	tests := []struct {
		in   []byte
		want string
	}{
		{[]byte{0x4f}, "4f"},
		{[]byte{0x0a}, "0a"},
		{[]byte{0xde, 0xad, 0xbe, 0xef}, "deadbeef"},
	}
	for _, tt := range tests {
		got := HexEncode(tt.in)
		if got != tt.want {
			t.Errorf("hexEncode(%x) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
