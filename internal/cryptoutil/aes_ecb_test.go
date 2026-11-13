package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestDecryptEcb(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	want := []byte("HELLO CRYPTOPALS")

	cipher, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	encrypted := make([]byte, 16)
	cipher.Encrypt(encrypted, want)

	got, err := DecryptEcb(encrypted, key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("DecryptEcb() = %q, want %q", got, want)
	}
}

func TestDetectEcb(t *testing.T) {
	input := [][]byte{
		[]byte("0123456789abcdeffedcba9876543210"),
		[]byte("YELLOW SUBMARINEYELLOW SUBMARINE"),
	}
	want := 1

	got, err := DetectEcb(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("DetectEcb() = %d, want %d", got, want)
	}
}
