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

func DetectEcb(input [][]byte) (int, error) {
	for i, line := range input {
		if len(line)%16 != 0 {
			continue
		}

		// Reiniciamos el mapa para CADA línea
		foundGroups := make(map[string]int)

		for j := 0; j+16 <= len(line); j += 16 {
			block := string(line[j : j+16])
			foundGroups[block]++

			// Si la cuenta llega a 2 (es decir, > 1), ya encontramos un duplicado
			if foundGroups[block] > 1 {
				return i, nil
			}
		}
	}
	return 0, errors.New("Ecb not found")
}
