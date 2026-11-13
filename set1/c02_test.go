package set1

import (
	"cryptopals/internal/cryptoutil"
	"testing"
)

func fixedStringXor(a string, b string) (string, error) {
	aBytes, aErr := cryptoutil.HexDecode(a)

	if aErr != nil {
		return "", aErr
	}

	bBytes, bErr := cryptoutil.HexDecode(b)

	if bErr != nil {
		return "", bErr
	}

	result, err := cryptoutil.FixedXor(aBytes, bBytes)

	if err != nil {
		return "", err
	}

	return cryptoutil.HexEncode(result), nil
}

func TestChallenge02_FixedStringXor(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"

	got, err := fixedStringXor(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
