package challenge4

import (
	"math"

	"cryptopals/internal/cryptoutil"
)

func findEncrypted(path string) (string, error) {
	strings, err := cryptoutil.ReadFileByLines(path)
	if err != nil {
		return "", err
	}

	bestScore := math.Inf(-1)
	bestResult := ""

	for _, str := range strings {
		decoded, e := cryptoutil.HexDecode(str)
		if e != nil {
			return "", e
		}
		c := cryptoutil.FindXorKey(decoded)

		if c.Score > bestScore {
			bestScore = c.Score
			bestResult = string(c.DecryptedString)
		}
	}

	return bestResult, nil
}
