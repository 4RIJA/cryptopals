package set1

import (
	"cryptopals/internal/cryptoutil"
	"testing"
)

func decryptXor(s string) (string, error) {
	rawBytes, err := cryptoutil.HexDecode(s)
	if err != nil {
		return "", err
	}
	c := cryptoutil.FindXorKey(rawBytes)
	return string(c.DecryptedString), nil
}

func TestChallenge03_DecryptXor(t *testing.T) {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"

	got, err := decryptXor(in)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
