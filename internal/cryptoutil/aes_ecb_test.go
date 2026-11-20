package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestEcbDecrypt(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	want := []byte("HELLO CRYPTOPALS")

	cipher, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	encrypted := make([]byte, 16)
	cipher.Encrypt(encrypted, want)

	got, err := EcbDecrypt(encrypted, key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("EcbDecrypt() = %q, want %q", got, want)
	}

	// Test invalid ciphertext length
	if _, err := EcbDecrypt([]byte("short"), key); err == nil {
		t.Errorf("expected error for non-multiple-of-16 ciphertext length, got nil")
	}

	// Test invalid key
	if _, err := EcbDecrypt(encrypted, []byte("short_key")); err == nil {
		t.Errorf("expected error for invalid key length, got nil")
	}
}

func TestEcbEncrypt(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")

	t.Run("encrypts and decrypts matching padded input", func(t *testing.T) {
		plaintext := []byte("YELLOW SUBMARINE")
		encrypted, err := EcbEncrypt(plaintext, key)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// A 16-byte input must be padded to 32 bytes with 16 bytes of 0x10
		if len(encrypted) != 32 {
			t.Fatalf("len(encrypted) = %d, want 32", len(encrypted))
		}

		decrypted, err := EcbDecrypt(encrypted, key)
		if err != nil {
			t.Fatalf("unexpected error decrypting: %v", err)
		}

		expectedPadded := PkCS7Padding(plaintext, 16)
		if !bytes.Equal(decrypted, expectedPadded) {
			t.Errorf("decrypted = %q, want %q", decrypted, expectedPadded)
		}
	})

	t.Run("encrypts partial block input", func(t *testing.T) {
		plaintext := []byte("HELLO WORLD")
		encrypted, err := EcbEncrypt(plaintext, key)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(encrypted) != 16 {
			t.Fatalf("len(encrypted) = %d, want 16", len(encrypted))
		}

		decrypted, err := EcbDecrypt(encrypted, key)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedPadded := PkCS7Padding(plaintext, 16)
		if !bytes.Equal(decrypted, expectedPadded) {
			t.Errorf("decrypted = %q, want %q", decrypted, expectedPadded)
		}
	})

	t.Run("matches stdlib aes cipher output", func(t *testing.T) {
		plaintext := []byte("YELLOW SUBMARINE")
		padded := PkCS7Padding(plaintext, 16)

		cipher, err := aes.NewCipher(key)
		if err != nil {
			t.Fatal(err)
		}

		expectedCiphertext := make([]byte, len(padded))
		for i := 0; i < len(padded); i += 16 {
			cipher.Encrypt(expectedCiphertext[i:i+16], padded[i:i+16])
		}

		gotCiphertext, err := EcbEncrypt(plaintext, key)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !bytes.Equal(gotCiphertext, expectedCiphertext) {
			t.Errorf("got %x, want %x", gotCiphertext, expectedCiphertext)
		}
	})

	t.Run("returns error on invalid key length", func(t *testing.T) {
		_, err := EcbEncrypt([]byte("data"), []byte("invalid"))
		if err == nil {
			t.Errorf("expected error for invalid key, got nil")
		}
	})
}

func TestDetectEcb(t *testing.T) {
	input := [][]byte{
		[]byte("0123456789abcdeffedcba9876543210"),
		[]byte("YELLOW SUBMARINEYELLOW SUBMARINE"),
	}
	want := 1

	got, err := DetectEcb(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("DetectEcb() = %d, want %d", got, want)
	}
}
