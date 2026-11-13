package set2

import (
	"bytes"
	"cryptopals/internal/cryptoutil"
	"testing"
)

func TestChallenge09_PkCS7(t *testing.T) {
	want := "YELLOW SUBMARINE\x04\x04\x04\x04"
	in := "YELLOW SUBMARINE"
	inBytes := []byte(in)

	got := cryptoutil.PkCS7Padding(inBytes, 20)

	if !bytes.Equal(got, []byte(want)) {
		t.Errorf("PkCS7Padding() = %q, want %q", got, want)
	}
}
