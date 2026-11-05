package cryptoutil

import (
	"testing"
)

func TestBinToBase64(t *testing.T) {
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
		got := binToBase64(tt.idx)
		if got != tt.want {
			t.Errorf("binToBase64(%d) = %q, want %q", tt.idx, got, tt.want)
		}
	}
}

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{"empty", []byte(""), ""},
		{"one remaining byte", []byte("M"), "TQ=="},
		{"two remaining bytes", []byte("Ma"), "TWE="},
		{"multiple of three", []byte("Man"), "TWFu"},
		{"one remaining byte with high bits", []byte{0xD3}, "0w=="},
		{"long phrase", []byte("Many hands make light work."), "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"},
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

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one decoded byte", "TQ==", "M"},
		{"two decoded bytes", "TWE=", "Ma"},
		{"multiple of three", "TWFu", "Man"},
		{"long phrase", "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu", "Many hands make light work."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Decode([]byte(tt.in))
			if err != nil {
				t.Fatalf("Base64Decode(%q) returned error: %v", tt.in, err)
			}
			if string(got) != tt.want {
				t.Errorf("Base64Decode(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
