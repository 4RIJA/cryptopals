package set2

import (
	"bytes"
	"cryptopals/internal/cryptoutil"
	"testing"
)

func TestChallenge11_ModeDetection(t *testing.T) {
	// A payload of at least 43 bytes of identical bytes is required so that regardless of
	// the random prefix (5-10 bytes) and suffix (5-10 bytes), at least two full 16-byte blocks
	// consist purely of the repeated plaintext under ECB. We provide 64 identical bytes.
	repeatedData := bytes.Repeat([]byte("A"), 64)

	var ecbCount, cbcCount int
	const iterations = 100

	for i := 0; i < iterations; i++ {
		ciphertext, actualMode, err := cryptoutil.EncryptionOracle(repeatedData)
		if err != nil {
			t.Fatalf("iteration %d: EncryptionOracle failed: %v", i, err)
		}

		detectedMode, err := cryptoutil.DetectMode(ciphertext)
		if err != nil {
			t.Fatalf("iteration %d: DetectMode failed: %v", i, err)
		}

		if detectedMode != actualMode {
			t.Fatalf("iteration %d: misdetection: got %v, want %v", i, detectedMode, actualMode)
		}

		switch actualMode {
		case cryptoutil.ModeECB:
			ecbCount++
		case cryptoutil.ModeCBC:
			cbcCount++
		}
	}

	if ecbCount == 0 || cbcCount == 0 {
		t.Errorf("expected both ECB and CBC to be tested, got ECB=%d, CBC=%d", ecbCount, cbcCount)
	}
}
