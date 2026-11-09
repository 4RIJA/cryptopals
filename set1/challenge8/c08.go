package challenge8

import (
	"cryptopals/internal/cryptoutil"
	"errors"
	"log"
)

// Devuelvo el indice de los bytes encriptados con aes ecb mode
func detectEcb(input [][]byte) (int, error) {
	for i, line := range input {
		if len(line)%16 != 0 {
			continue
		}

		// Reiniciamos el mapa para CADA línea
		foundGroups := make(map[string]int)

		for j := 0; j+16 <= len(line); j += 16 {
			block := string(line[j : j+16])
			foundGroups[block]++

			// Si la cuenta llega a 2 (es decir, > 1), ya encontramos un duplicado
			if foundGroups[block] > 1 {
				return i, nil
			}
		}
	}
	return 0, errors.New("Ecb not found")
}

func Challenge8() int {
	lines, err := cryptoutil.ReadFileByLines("set1/challenge8/8.txt")
	if err != nil {
		log.Fatal(err)
	}
	decodedLines := make([][]byte, 0, len(lines))

	for _, l := range lines {
		decodedLine, err := cryptoutil.HexDecode(l)
		if err != nil {
			log.Fatal(err)
		}
		decodedLines = append(decodedLines, decodedLine)
	}

	ecbIndex, err := detectEcb(decodedLines)
	if err != nil {
		log.Fatal(err)
	}
	return ecbIndex
}
