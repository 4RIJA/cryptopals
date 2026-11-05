package challenge6

import (
	"cmp"
	"cryptopals/internal/cryptoutil"
	"errors"
	"math"
	"slices"
)

func hammingDistance(aBytes []byte, bBytes []byte) (int, error) {
	if len(aBytes) != len(bBytes) {
		return 0, errors.New("Both strings must be the same length")
	}
	distance := 0

	xorBytes, err := cryptoutil.FixedXor(aBytes, bBytes)
	if err != nil {
		return 0, err
	}

	for _, b := range xorBytes {
		distance += kenighanAlgorithm(b)
	}
	return distance, nil

}

func BreakRepeatingKeyXor(input []byte) []byte {
	var finalKey []byte
	maxScore := math.Inf(-1)

	keySizes := findMostProbableKeySizes(input, 2, 40, 3)

	for _, ks := range keySizes {
		k := keyFromSize(input, ks)
		deCipher := cryptoutil.XorRepeatedKeys(k, input)
		score := cryptoutil.ScoreEnglish(deCipher)
		if score > maxScore {
			finalKey = k
			maxScore = score
		}
	}

	return finalKey
}

// Devuelve la key que mejor resultado devuelve
func keyFromSize(input []byte, size int) []byte {
	var key []byte
	m := cryptoutil.SplitIntoBlocks(input, size)
	transposedM := cryptoutil.Transpose(m)
	for _, block := range transposedM {
		c := cryptoutil.FindXorKey(block)
		key = append(key, c.Key)
	}
	return key
}

type KeySize struct {
	size        int
	hammingDist float64
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
			distance, _ := hammingDistance(group1, group2)
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

// Este algoritmo calcula el numero de bits a 1 eliminando de derecha a izquierda los 1 existentes en el byte. Esto pasa cuando comparas un byte con el anterior.
func kenighanAlgorithm(b byte) int {
	count := 0
	for b != 0x00 {
		count++
		b &= b - 1
	}
	return count
}
