package challenge9

import (
	"bytes"
	"testing"
)

func TestPkCS7(t *testing.T) {
	want := "YELLOW SUBMARINE\x04\x04\x04\x04"
	in := "YELLOW SUBMARINE"
	inBytes := []byte(in)

	got := pkCS7Padding(inBytes, 20)

	if !bytes.Equal(got, []byte(want)) {
		t.Errorf("pkCS7Padding() = %q, want %q", got, want)
	}
}
