package cryptoutil

import (
	"reflect"
	"testing"
)

func TestSplitIntoBlocks(t *testing.T) {
	in := []byte{1, 2, 3, 4, 5}
	want := [][]byte{
		{1, 2},
		{3, 4},
		{5},
	}

	got := SplitIntoBlocks(in, 2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SplitIntoBlocks(%v, 2) = %v, want %v", in, got, want)
	}
}

func TestTranspose(t *testing.T) {
	in := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := [][]byte{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	got := Transpose(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transpose(%v) = %v, want %v", in, got, want)
	}
}
