package cryptoutil

import (
	"math"
	"testing"
)

func TestScoreEnglishReturnsAverageScore(t *testing.T) {
	singleByteScore := ScoreEnglish([]byte("e"))
	repeatedByteScore := ScoreEnglish([]byte("eeee"))
	if repeatedByteScore != singleByteScore {
		t.Errorf("ScoreEnglish(%q) = %v, want %v", "eeee", repeatedByteScore, singleByteScore)
	}
}

func TestScoreEnglishEmptyInput(t *testing.T) {
	if score := ScoreEnglish(nil); !math.IsInf(score, -1) {
		t.Errorf("ScoreEnglish(nil) = %v, want negative infinity", score)
	}
}
