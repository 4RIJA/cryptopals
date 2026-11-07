package challenge7

import (
	"crypto/aes"
	"cryptopals/internal/cryptoutil"
	"errors"
	"log"
	"strings"
)

/*
Si nos fijamos en el resultado, al final del texto descrifrado hay 4bytes a 0x04,
esto es porque el emisor a la hora de cifrar el mansaje en aes-128 ecb mode tuvo
que añadir ese pad para que el mensaje fuese multiplo de 16.
A esta regla se le llama PKCS#7.
*/

func decryptEcb(input []byte, key []byte) ([]byte, error) {
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

func Challenge7() string {
	lines, err := cryptoutil.ReadFileByLines("7.txt")
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	base64String := strings.Join(lines, "")

	decodedString, err := cryptoutil.Base64Decode([]byte(base64String))
	if err != nil {
		log.Fatalf("Error decoding")
	}

	decryptedString, err := decryptEcb(decodedString, []byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return string(decryptedString)
}
