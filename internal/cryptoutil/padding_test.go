package cryptoutil

import (
	"bytes"
	"testing"
)

func TestPkCS7Padding(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		size int
		want []byte
	}{
		{"pad 4 bytes", []byte("YELLOW SUBMARINE"), 20, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")},
		{"exact multiple adds full block", []byte("ABCD"), 4, []byte("ABCD\x04\x04\x04\x04")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PkCS7Padding(tt.in, tt.size)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("PkCS7Padding(%q, %d) = %q, want %q", tt.in, tt.size, got, tt.want)
			}
		})
	}
}
