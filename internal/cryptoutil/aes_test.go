package cryptoutil

import (
	"bytes"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	t.Run("generates requested size", func(t *testing.T) {
		sizes := []int{0, 1, 16, 24, 32, 64}
		for _, size := range sizes {
			rnd := GenerateRandomBytes(size)
			if len(rnd) != size {
				t.Errorf("GenerateRandomBytes(%d) returned slice of length %d", size, len(rnd))
			}
		}
	})

	t.Run("generates random different bytes on successive calls", func(t *testing.T) {
		b1 := GenerateRandomBytes(16)
		b2 := GenerateRandomBytes(16)

		if bytes.Equal(b1, b2) {
			t.Errorf("GenerateRandomBytes(16) returned identical slices: %x == %x", b1, b2)
		}
	})
}

func TestEncryptionOracle(t *testing.T) {
	t.Run("returns valid ciphertext of multiple of 16 bytes", func(t *testing.T) {
		inputs := [][]byte{
			{},
			[]byte("A"),
			[]byte("YELLOW SUBMARINE"),
			bytes.Repeat([]byte("X"), 64),
		}

		for _, in := range inputs {
			encrypted, mode, err := EncryptionOracle(in)
			if err != nil {
				t.Fatalf("EncryptionOracle returned error: %v", err)
			}
			if len(encrypted) == 0 {
				t.Fatalf("EncryptionOracle returned empty ciphertext")
			}
			if len(encrypted)%16 != 0 {
				t.Errorf("ciphertext length %d is not a multiple of 16", len(encrypted))
			}
			if mode != ModeECB && mode != ModeCBC {
				t.Errorf("unexpected mode returned: %v", mode)
			}
		}
	})

	t.Run("consistently succeeds and produces varying outputs over multiple runs", func(t *testing.T) {
		seen := make(map[string]bool)
		input := []byte("constant input data for testing")

		for i := 0; i < 50; i++ {
			encrypted, _, err := EncryptionOracle(input)
			if err != nil {
				t.Fatalf("iteration %d: unexpected error %v", i, err)
			}
			if len(encrypted)%16 != 0 {
				t.Fatalf("iteration %d: length %d not a multiple of 16", i, len(encrypted))
			}
			seen[string(encrypted)] = true
		}

		// Because random keys, IVs, prefixes, and suffixes are chosen, every output should be unique
		if len(seen) != 50 {
			t.Errorf("expected 50 unique ciphertexts, got %d", len(seen))
		}
	})

	t.Run("exercises both ECB and CBC modes", func(t *testing.T) {
		input := bytes.Repeat([]byte("A"), 64)

		foundECB := false
		foundCBC := false

		for i := 0; i < 100; i++ {
			_, mode, err := EncryptionOracle(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if mode == ModeECB {
				foundECB = true
			} else if mode == ModeCBC {
				foundCBC = true
			}

			if foundECB && foundCBC {
				break
			}
		}

		if !foundECB {
			t.Errorf("expected to observe ECB mode across 100 runs")
		}
		if !foundCBC {
			t.Errorf("expected to observe CBC mode across 100 runs")
		}
	})
}

func TestDetectMode(t *testing.T) {
	t.Run("detects ECB when duplicate blocks exist", func(t *testing.T) {
		block := []byte("0123456789abcdef")
		// 3 identical blocks
		ciphertext := bytes.Repeat(block, 3)

		mode, err := DetectMode(ciphertext)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if mode != ModeECB {
			t.Errorf("DetectMode() = %v, want %v", mode, ModeECB)
		}
	})

	t.Run("detects CBC when no duplicate blocks exist", func(t *testing.T) {
		b1 := []byte("0123456789abcdef")
		b2 := []byte("fedcba9876543210")
		b3 := []byte("abcdef0123456789")
		ciphertext := append(append(b1, b2...), b3...)

		mode, err := DetectMode(ciphertext)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if mode != ModeCBC {
			t.Errorf("DetectMode() = %v, want %v", mode, ModeCBC)
		}
	})

	t.Run("handles empty or short input gracefully", func(t *testing.T) {
		mode, err := DetectMode([]byte{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if mode != ModeCBC {
			t.Errorf("DetectMode([]byte{}) = %v, want %v", mode, ModeCBC)
		}
	})
}
