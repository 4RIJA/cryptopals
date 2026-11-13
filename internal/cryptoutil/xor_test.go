package cryptoutil

import (
	"bytes"
	"testing"
)

func TestXor(t *testing.T) {
	if got := xor(0x1c, 0x68); got != 0x74 {
		t.Errorf("xor(0x1c, 0x68) = %#x, want 0x74", got)
	}
	if got := xor(0xff, 0xff); got != 0x00 {
		t.Errorf("xor(0xff, 0xff) = %#x, want 0x00", got)
	}
}

func TestFixedXor(t *testing.T) {
	a := []byte{0x1c, 0x01, 0x11, 0x00}
	b := []byte{0x68, 0x69, 0x74, 0x20}
	want := []byte{0x74, 0x68, 0x65, 0x20}

	got, err := FixedXor(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("FixedXor(%x, %x) = %x, want %x", a, b, got, want)
	}
}

func TestXorRepeatedKeys(t *testing.T) {
	key := []byte("ICE")
	in := []byte("Burning 'em")
	want := []byte{0x0b, 0x36, 0x37, 0x27, 0x2a, 0x2b, 0x2e, 0x63, 0x62, 0x2c, 0x2e}

	got := XorRepeatedKeys(key, in)
	if !bytes.Equal(got, want) {
		t.Errorf("XorRepeatedKeys(%q, %q) = %x, want %x", key, in, got, want)
	}
}

func TestFindXorKey(t *testing.T) {
	wantKey := byte('X')
	wantText := "Cooking MC's like a pound of bacon"
	encrypted := XorRepeatedKeys([]byte{wantKey}, []byte(wantText))

	got := FindXorKey(encrypted)
	if got.Key != wantKey {
		t.Errorf("FindXorKey() key = %q, want %q", got.Key, wantKey)
	}
	if string(got.DecryptedString) != wantText {
		t.Errorf("FindXorKey() text = %q, want %q", got.DecryptedString, wantText)
	}
}

func TestCountSetBits(t *testing.T) {
	tests := []struct {
		in   byte
		want int
	}{
		{0x00, 0},
		{0x09, 2},
		{0x07, 3},
		{0xff, 8},
	}

	for _, tt := range tests {
		got := kernighanAlgorithm(tt.in)
		if got != tt.want {
			t.Errorf("kernighanAlgorithm(%#x) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	in1 := "this is a test"
	in2 := "wokka wokka!!!"
	want := 37

	aBytes := []byte(in1)
	bBytes := []byte(in2)

	got, err := HammingDistance(aBytes, bBytes)

	if err != nil {
		t.Fatalf("HammingDistance returned an error: %v", err)
	}

	if got != want {
		t.Errorf("HammingDistance(%q, %q) = %d, want %d", in1, in2, got, want)
	}
}
