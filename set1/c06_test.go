package set1

import (
	"cryptopals/internal/cryptoutil"
	"strings"
	"testing"
)

func TestChallenge06_BreakRepeatingKeyXor(t *testing.T) {
	lines, err := cryptoutil.ReadFileByLines("testdata/6.txt")
	if err != nil {
		t.Fatalf("could not read challenge input: %v", err)
	}
	input, err := cryptoutil.Base64Decode([]byte(strings.Join(lines, "")))
	if err != nil {
		t.Fatalf("could not decode challenge input: %v", err)
	}

	got := cryptoutil.BreakRepeatingKeyXor(input)
	const want = "Terminator X: Bring the noise"
	if string(got) != want {
		t.Errorf("BreakRepeatingKeyXor() = %q, want %q", got, want)
	}
}
