package challenge1

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
			got, err := hexDecode(tt.in)
			if err != nil {
				t.Fatalf("hexDecode(%q) devolvio error: %v", tt.in, err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("hexDecode(%q) = %x, want %x", tt.in, got, tt.want)
			}
		})
	}
}

func TestBase64Value(t *testing.T) {
	tests := []struct {
		idx  byte
		want byte
	}{
		{0, 'A'},
		{10, 'K'},
		{25, 'Z'},
		{26, 'a'},
		{36, 'k'},
		{51, 'z'},
		{52, '0'},
		{61, '9'},
		{62, '+'},
		{63, '/'},
	}
	for _, tt := range tests {
		got := base64Value(tt.idx)
		if got != tt.want {
			t.Errorf("base64Value(%d) = %q, want %q", tt.idx, got, tt.want)
		}
	}
}

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{"vacio", []byte(""), ""},
		{"sobra 1 byte", []byte("M"), "TQ=="},
		{"sobran 2 bytes", []byte("Ma"), "TWE="},
		{"multiplo de 3", []byte("Man"), "TWFu"},
		{"sobra 1 byte con bits altos", []byte{0xD3}, "0w=="},
		{"frase larga", []byte("Many hands make light work."), "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base64Encode(tt.in)
			if got != tt.want {
				t.Errorf("base64Encode(%x) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestChallenge01_HexToBase64(t *testing.T) {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	got, err := hexToBase64(in)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
