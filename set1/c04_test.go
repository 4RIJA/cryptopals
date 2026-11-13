package set1

import (
	"cryptopals/internal/cryptoutil"
	"math"
	"strings"
	"testing"
)

func findEncrypted(path string) (string, error) {
	lines, err := cryptoutil.ReadFileByLines(path)
	if err != nil {
		return "", err
	}

	bestScore := math.Inf(-1)
	bestResult := ""

	for _, str := range lines {
		decoded, e := cryptoutil.HexDecode(str)
		if e != nil {
			return "", e
		}
		c := cryptoutil.FindXorKey(decoded)

		if c.Score > bestScore {
			bestScore = c.Score
			bestResult = string(c.DecryptedString)
		}
	}

	return bestResult, nil
}

func TestChallenge04_FindEncrypted(t *testing.T) {
	want := "Now that the party is jumping"

	got, err := findEncrypted("testdata/4.txt")
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
