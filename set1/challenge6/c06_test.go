package challenge6

import (
	"cryptopals/internal/cryptoutil"
	"strings"
	"testing"
)

func TestChallenge06_BreakRepeatingKeyXor(t *testing.T) {
	lines, err := cryptoutil.ReadFileByLines("6.txt")
	if err != nil {
		t.Fatalf("could not read challenge input: %v", err)
	}
	input, err := cryptoutil.Base64Decode([]byte(strings.Join(lines, "")))
	if err != nil {
		t.Fatalf("could not decode challenge input: %v", err)
	}

	got := BreakRepeatingKeyXor(input)
	const want = "Terminator X: Bring the noise"
	if string(got) != want {
		t.Errorf("BreakRepeatingKeyXor() = %q, want %q", got, want)
	}
}

func TestCountSetBits(t *testing.T) {
	tests := []struct {
		in   byte
		want int
	}{
		{0x00, 0},
		{0x09, 2},
		{0x07, 3},
		{0xff, 8},
	}

	for _, tt := range tests {
		got := kernighanAlgorithm(tt.in)
		if got != tt.want {
			t.Errorf("kernighanAlgorithm(%#x) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	in1 := "this is a test"
	in2 := "wokka wokka!!!"
	want := 37

	aBytes := []byte(in1)
	bBytes := []byte(in2)

	got, err := hammingDistance(aBytes, bBytes)

	if err != nil {
		t.Fatalf("hammingDistance returned an error: %v", err)
	}

	if got != want {
		t.Errorf("hamminDistance(%q, %q) = %d, want %d", in1, in2, got, want)
	}
}
