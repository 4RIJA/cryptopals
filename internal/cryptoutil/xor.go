package cryptoutil

import (
	"cmp"
	"errors"
	"math"
	"slices"
)

type Candidate struct {
	DecryptedString []byte
	Key             byte
	Score           float64
}

type KeySize struct {
	size        int
	hammingDist float64
}

func FixedXor(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("not the same length")
	}

	result := make([]byte, len(b))
	for i := range a {
		result[i] = xor(a[i], b[i])
	}
	return result, nil
}

func XorRepeatedKeys(keys []byte, toEncrypt []byte) []byte {
	result := make([]byte, len(toEncrypt))
	for i, v := range toEncrypt {
		key := keys[i%len(keys)]
		result[i] = xor(v, key)
	}
	return result
}

func xor(a byte, b byte) byte {
	return a ^ b
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

func HammingDistance(aBytes []byte, bBytes []byte) (int, error) {
	if len(aBytes) != len(bBytes) {
		return 0, errors.New("Both strings must be the same length")
	}
	distance := 0

	xorBytes, err := FixedXor(aBytes, bBytes)
	if err != nil {
		return 0, err
	}

	for _, b := range xorBytes {
		distance += kernighanAlgorithm(b)
	}
	return distance, nil
}

func BreakRepeatingKeyXor(input []byte) []byte {
	var finalKey []byte
	maxScore := math.Inf(-1)

	keySizes := findMostProbableKeySizes(input, 2, 40, 3)

	for _, ks := range keySizes {
		k := keyFromSize(input, ks)
		deCipher := XorRepeatedKeys(k, input)
		score := ScoreEnglish(deCipher)
		if score > maxScore {
			finalKey = k
			maxScore = score
		}
	}

	return finalKey
}

// Devuelve la clave que obtiene la mejor puntuación.
func keyFromSize(input []byte, size int) []byte {
	var key []byte
	m := SplitIntoBlocks(input, size)
	transposedM := Transpose(m)
	for _, block := range transposedM {
		c := FindXorKey(block)
		key = append(key, c.Key)
	}
	return key
}

func findMostProbableKeySizes(input []byte, from, to, numberOfKeys int) []int {
	keySizes := make([]KeySize, 0, to-from+1)

	for keySize := from; keySize <= to; keySize++ {
		// Necesitamos al menos dos bloques completos
		if 2*keySize > len(input) {
			break
		}

		numBlocks := len(input) / keySize

		// Media de la distancia entre bloques consecutivos
		sumDistances := 0
		for j := 1; j < numBlocks; j++ {
			group1 := input[(j-1)*keySize : j*keySize]
			group2 := input[j*keySize : (j+1)*keySize]
			distance, _ := HammingDistance(group1, group2)
			sumDistances += distance
		}

		// Normalizamos por keySize para poder comparar entre tamaños distintos
		averageDist := (float64(sumDistances) / float64(numBlocks-1)) / float64(keySize)
		keySizes = append(keySizes, KeySize{keySize, averageDist})
	}

	// Ordenamos
	slices.SortFunc(keySizes, func(a, b KeySize) int {
		return cmp.Compare(a.hammingDist, b.hammingDist)
	})

	if numberOfKeys > len(keySizes) {
		numberOfKeys = len(keySizes)
	}

	sizes := make([]int, 0, numberOfKeys)
	for _, ks := range keySizes[:numberOfKeys] {
		sizes = append(sizes, ks.size)
	}

	return sizes
}

// Cuenta los bits a uno eliminando, en cada iteración, el bit activo situado más a la derecha.
func kernighanAlgorithm(b byte) int {
	count := 0
	for b != 0x00 {
		count++
		b &= b - 1
	}
	return count
}
