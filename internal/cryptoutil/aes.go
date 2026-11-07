package cryptoutil

import (
	"crypto/aes"
	"errors"
)

func DecryptEcb(input []byte, key []byte) ([]byte, error) {
	if len(input)%16 != 0 {
		return nil, errors.New("The cipher text is not valid for decryption. The length isnt multiple of 16.")
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(input))
	for i := 0; i+16 <= len(input); i += 16 {
		cipher.Decrypt(result[i:min(i+16, len(result))], input[i:min(i+16, len(input))])
	}
	return result, nil
}
