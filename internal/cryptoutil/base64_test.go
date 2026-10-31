package cryptoutil

import (
	"testing"
)

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
			got := Base64Encode(tt.in)
			if got != tt.want {
				t.Errorf("base64Encode(%x) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
