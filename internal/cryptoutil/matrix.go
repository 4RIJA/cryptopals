package cryptoutil

func SplitIntoBlocks(data []byte, n int) [][]byte {
	blocks := make([][]byte, 0, (len(data)+n-1)/n)
	for i := 0; i < len(data); i += n {
		end := min(i+n, len(data))
		blocks = append(blocks, data[i:end])
	}
	return blocks
}

// Transpose transpone una matriz.
func Transpose[T any](m [][]T) [][]T {
	transposedMatrix := make([][]T, len(m[0]))
	for _, row := range m {
		for j := range row {
			transposedMatrix[j] = append(transposedMatrix[j], row[j])
		}
	}
	return transposedMatrix
}
