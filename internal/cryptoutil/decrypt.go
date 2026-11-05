package cryptoutil

import "math"

type Candidate struct {
	DecryptedString []byte
	Key             byte
	Score           float64
}

func FindXorKey(s []byte) Candidate {
	bestScore := math.Inf(-1)
	var bestKey byte
	var maxResult []byte
	for i := range 256 {
		b := XorRepeatedKeys([]byte{byte(i)}, s)
		currentScore := ScoreEnglish(b)
		if currentScore > bestScore {
			bestScore = currentScore
			bestKey = byte(i)
			maxResult = b
		}
	}
	return Candidate{DecryptedString: maxResult, Key: bestKey, Score: bestScore}
}
