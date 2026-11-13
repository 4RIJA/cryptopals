package set1

import (
	"cryptopals/internal/cryptoutil"
	"testing"
)

func TestChallenge08_DetectEcb(t *testing.T) {
	lines, err := cryptoutil.ReadFileByLines("testdata/8.txt")
	if err != nil {
		t.Fatal(err)
	}
	decodedLines := make([][]byte, 0, len(lines))

	for _, l := range lines {
		decodedLine, err := cryptoutil.HexDecode(l)
		if err != nil {
			t.Fatal(err)
		}
		decodedLines = append(decodedLines, decodedLine)
	}

	got, err := cryptoutil.DetectEcb(decodedLines)
	if err != nil {
		t.Fatal(err)
	}

	const want = 132
	if got != want {
		t.Errorf("DetectEcb() = %d, want %d", got, want)
	}
}
