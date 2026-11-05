package challenge3

import (
	"math"

	"cryptopals/internal/cryptoutil"
)

func decryptXor(s string) (string, error) {
	rawBytes, err := cryptoutil.HexDecode(s)
	if err != nil {
		return "", err
	}
	bestScore := math.Inf(-1)
	var maxResult []byte
	for i := range 256 {
		b := cryptoutil.XorRepeatedKeys([]byte{byte(i)}, rawBytes)
		currentScore := cryptoutil.ScoreEnglish(b)
		if currentScore > bestScore {
			bestScore = currentScore
			maxResult = b
		}
	}
	return string(maxResult), nil
}
