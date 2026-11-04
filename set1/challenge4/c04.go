package challenge4

import "cryptopals/internal/cryptoutil"

func findEncrypted(path string) (string, error) {
	strings, err := cryptoutil.ReadFileByLines(path)
	if err != nil {
		return "", err
	}

	bestScore := 0.0
	bestResult := ""

	for _, str := range strings {
		c, e := cryptoutil.DecryptXor(str)

		if e != nil {
			return "", e
		}

		if c.Score > bestScore {
			bestScore = c.Score
			bestResult = string(c.DecryptedString)
		}
	}

	return bestResult, nil
}
