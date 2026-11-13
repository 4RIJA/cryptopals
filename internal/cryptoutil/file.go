package cryptoutil

import (
	"bufio"
	"os"
)

func ReadFileByLines(path string) ([]string, error) {
	var results []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
