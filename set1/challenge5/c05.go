package challenge5

import (
	"cryptopals/internal/cryptoutil"
)

func cipherXorWithICE(text string) string {
	encryptedBytes := cryptoutil.XorRepeatedKeys([]byte{'I', 'C', 'E'}, []byte(text))
	resultHex := cryptoutil.HexEncode(encryptedBytes)
	return resultHex
}
