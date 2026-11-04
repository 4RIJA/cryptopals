package cryptoutil

type Candidate struct {
	DecryptedString []byte
	Key             byte
	Score           float64
}

func DecryptXor(s string) (Candidate, error) {
	rawBytes, err := HexDecode(s)
	if err != nil {
		return Candidate{}, err
	}
	bestScore := 0.0
	var bestKey byte
	var maxResult []byte
	for i := range 256 {
		b := XorRepeatedKeys([]byte{byte(i)}, rawBytes)
		currentScore := ScoreEnglish(b)
		if currentScore > bestScore {
			bestScore = currentScore
			bestKey = byte(i)
			maxResult = b
		}
	}
	return Candidate{DecryptedString: maxResult, Key: bestKey, Score: bestScore}, nil
}
