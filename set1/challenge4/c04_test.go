package challenge4

import (
	"strings"
	"testing"
)

func TestChallenge04_FindEncrypted(t *testing.T) {
	want := "Now that the party is jumping"

	got, err := findEncrypted("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
