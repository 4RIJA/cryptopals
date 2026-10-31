package challenge1

import (
	"cryptopals/internal/cryptoutil"
)

func hexToBase64(s string) (string, error) {
	bin, err := cryptoutil.HexDecode(s)
	if err != nil {
		return "", err
	}
	return cryptoutil.Base64Encode(bin), nil
}
