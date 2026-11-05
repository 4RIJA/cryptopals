package challenge2

import (
	"cryptopals/internal/cryptoutil"
)

func fixedStringXor(a string, b string) (string, error) {
	aBytes, aErr := cryptoutil.HexDecode(a)

	if aErr != nil {
		return "", aErr
	}

	bBytes, bErr := cryptoutil.HexDecode(b)

	if bErr != nil {
		return "", bErr
	}

	result, err := cryptoutil.FixedXor(aBytes, bBytes)

	if err != nil {
		return "", err
	}

	return cryptoutil.HexEncode(result), nil
}
