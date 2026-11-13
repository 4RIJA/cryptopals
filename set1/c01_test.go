package set1

import (
	"cryptopals/internal/cryptoutil"
	"testing"
)

func hexToBase64(s string) (string, error) {
	bin, err := cryptoutil.HexDecode(s)
	if err != nil {
		return "", err
	}
	return cryptoutil.Base64Encode(bin), nil
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
