package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"errors"
	"slices"
)

/*
Si nos fijamos en el resultado, al final del texto descrifrado hay 4bytes a 0x04,
esto es porque el emisor a la hora de cifrar el mansaje en aes-128 ecb mode tuvo
que añadir ese pad para que el mensaje fuese multiplo de 16.
A esta regla se le llama PKCS#7.
*/

func EcbDecrypt(input []byte, key []byte) ([]byte, error) {
	if len(input)%16 != 0 {
		return nil, errors.New("The cipher text is not valid for decryption. The length isnt multiple of 16.")
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(input))
	for i := 0; i+16 <= len(input); i += 16 {
		cipher.Decrypt(result[i:i+16], input[i:i+16])
	}
	return result, nil
}

func EcbEncrypt(input []byte, key []byte) ([]byte, error) {
	paddedInput := PkCS7Padding(input, 16)
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(paddedInput))
	for i := 0; i+16 <= len(paddedInput); i += 16 {
		cipher.Encrypt(result[i:i+16], paddedInput[i:i+16])
	}
	return result, nil
}

// Devuelvo el indice de los bytes encriptados con aes ecb mode
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

func CreateOracleECB(inputToDecrypt []byte) func([]byte) []byte {
	globalRandomKey := GenerateRandomBytes(16)

	return func(input []byte) []byte {
		encrypted, _ := EcbEncrypt(slices.Concat(input, inputToDecrypt), globalRandomKey)
		return encrypted
	}
}

func EcbDecryptByteATime(inputToDecrypt []byte) ([]byte, error) {
	ecbOracle := CreateOracleECB(inputToDecrypt)

	// Hay que averiguar el tamaño de los bloques para ello nos basaremos en el funcionamiento
	// del padding PKCS7. Cuando no hace falta paading porque la longitud del texto a encriptar
	// es multiplo de la de la clave. Lo abordaremos añadiendo bytes al input hasta que la funcion
	// del padding añada otro bloque, en ese punto compararemos la longitud inicial con la nueva
	// sacando asi la longitud de los bloques.
	initialEncrypted := ecbOracle([]byte{})
	initialLen := len(initialEncrypted)
	blockSize := 0
	bytesToJump := 0
	for i := range initialLen + 1 {
		encryptedResult := ecbOracle(bytes.Repeat([]byte{'A'}, i))
		if len(encryptedResult) > initialLen {
			blockSize = len(encryptedResult) - initialLen
			bytesToJump = i
			break
		}
	}
	dataSize := initialLen - bytesToJump

	// Comprobamos si el modo de encriptacion es ECB enviando bytes repetidos
	probeSize := 64
	if blockSize > 0 {
		probeSize = blockSize * 2
	}
	mode, err := DetectMode(ecbOracle(bytes.Repeat([]byte{'A'}, probeSize)))
	if err != nil {
		return nil, err
	}
	if mode != ModeECB {
		return nil, errors.New("El modo de encriptacion no es ecb")
	}

	// En el ejercicio nos piden que hagamos
	foundBytes := make([]byte, 0, dataSize)
	for i := 0; i+blockSize <= len(initialEncrypted); i += blockSize {
		blockStart := i
		blockEnd := i + blockSize

		for j := blockSize - 1; j >= 0; j-- {
			roundBytes := bytes.Repeat([]byte{'A'}, j)
			a := ecbOracle(roundBytes)[blockStart:blockEnd]
			for k := range 256 {
				b := ecbOracle(slices.Concat(roundBytes, foundBytes, []byte{byte(k)}))[blockStart:blockEnd]
				if bytes.Equal(a, b) {
					foundBytes = append(foundBytes, byte(k))
					break
				}
			}
		}
	}
	return foundBytes[:dataSize], nil
}
