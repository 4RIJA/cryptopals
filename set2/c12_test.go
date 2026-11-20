package set2

import (
	"bytes"
	"cryptopals/internal/cryptoutil"
	"testing"
)

func TestChallenge12_ByteAtATime(t *testing.T) {
	const secretB64 = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"

	want, err := cryptoutil.Base64Decode([]byte(secretB64))
	if err != nil {
		t.Fatal(err)
	}

	got, err := cryptoutil.EcbDecryptByteATime(want)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
