package cryptoutil

import "testing"

func TestXor(t *testing.T) {
	if got := xor(0x1c, 0x68); got != 0x74 {
		t.Errorf("xor(0x1c, 0x68) = %#x, want 0x74", got)
	}
	if got := xor(0xff, 0xff); got != 0x00 {
		t.Errorf("xor(0xff, 0xff) = %#x, want 0x00", got)
	}
}
