package cryptoutil

import (
	"crypto/aes"
	"errors"
)

func CbcEncrypt(input []byte, iv []byte, key []byte) ([]byte, error) {

	// Primero hay que mirar si es multiplo de 16 y en caso de que no lo sea hacer
	// padding PkCS7
	input = PkCS7Padding(input, 16)
	encryptedResult := make([]byte, 0, len(input))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	prevBlock := iv

	for i := 0; i+16 <= len(input); i += 16 {
		plainTextBlock := input[i:min(i+16, len(input))]
		xoredBlock, err := FixedXor(plainTextBlock, prevBlock)
		if err != nil {
			return nil, err
		}
		encryptedBlock := make([]byte, 16)
		cipher.Encrypt(encryptedBlock, xoredBlock)
		encryptedResult = append(encryptedResult, encryptedBlock...)
		prevBlock = encryptedBlock
	}

	return encryptedResult, nil
}

func CbcDecrypt(input []byte, iv []byte, key []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("No ciphered text provided")
	}
	if len(input)%16 != 0 {
		return nil, errors.New("the ciphered text isn't multiple of 16")
	}

	decryptedResult := make([]byte, 0, len(input))

	// Variable auxiliar para guardar el bloque con el que haremos xor tras desencriptar para
	// obtener el bloque descifrado. El primer bloque para hacer Xor es el iv.
	prevBlock := iv
	for i := 0; i+16 <= len(input); i += 16 {
		// Teniendo en cuenta que el metodo de aes para encriptar (al igual que para encriptar)
		// creamos una varibale donde el método guardará el bloque desencriptado
		decryptedBlock := make([]byte, 16)
		cipher, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		cipher.Decrypt(decryptedBlock, input[i:min(i+16, len(input))])

		// Hacemos el XOR para deshacer los XOR que se hacen el la encriptacion. Si es la primer
		// bloque se usará iv, sino se usara el el anterior bloque cifrado
		xoredBlock, err := FixedXor(decryptedBlock, prevBlock)
		if err != nil {
			return nil, err
		}
		// Ese bloque lo guardamos en el variable decryptedResult
		decryptedResult = append(decryptedResult, xoredBlock...)

		prevBlock = input[i:min(i+16, len(input))]
	}
	return decryptedResult, nil
}
