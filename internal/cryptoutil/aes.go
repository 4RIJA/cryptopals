package cryptoutil

import (
	"math/rand/v2"
)

type AESMode int

const (
	ModeECB AESMode = iota // 0
	ModeCBC                // 1
)

func EncryptionOracle(data []byte) ([]byte, AESMode, error) {
	mode := AESMode(rand.IntN(2))
	randomKey := GenerateRandomBytes(16)
	randomIv := GenerateRandomBytes(16)

	prefix := GenerateRandomBytes(rand.IntN(6) + 5)
	sufix := GenerateRandomBytes(rand.IntN(6) + 5)
	blockToEncrypt := make([]byte, 0, len(prefix)+len(data)+len(sufix))
	blockToEncrypt = append(blockToEncrypt, prefix...)
	blockToEncrypt = append(blockToEncrypt, data...)
	blockToEncrypt = append(blockToEncrypt, sufix...)

	if mode == ModeCBC {
		encrypted, err := CbcEncrypt(blockToEncrypt, randomIv, randomKey)
		if err != nil {
			return nil, mode, err
		}
		return encrypted, mode, nil
	} else {
		encrypted, err := EcbEncrypt(blockToEncrypt, randomKey)
		if err != nil {
			return nil, mode, err
		}
		return encrypted, mode, nil
	}
}

func GenerateRandomBytes(size int) []byte {
	rnd := make([]byte, size)
	for i := range size {
		rnd[i] = byte(rand.IntN(256))
	}
	return rnd
}

// Esta deteccion se usa principalmente en ataques activos donde, en este caso, es necesario enviar al menos
// 43 bytes iguales para que los paddings aleatorios no interfieran en la deteccion y se garanticen dos
// bloques exactamente iguales. Este tipo de ataque es cuando puedes por ejemplo enviar una imagen e interceptar
// el paquete y ver el archivo cifrado como se mueve por la red y asi determinar el metodo que se esta usando
func DetectMode(input []byte) (AESMode, error) {
	foundGroups := make(map[string]int)
	for i := 0; i+16 <= len(input); i += 16 {
		block := input[i : i+16]
		foundGroups[string(block)]++

		if foundGroups[string(block)] >= 2 {
			return ModeECB, nil
		}
	}
	return ModeCBC, nil
}
