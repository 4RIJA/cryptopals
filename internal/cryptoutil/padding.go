package cryptoutil

func PkCS7Padding(input []byte, size int) []byte {
	lenToPad := size - len(input)%size

	padded := make([]byte, len(input)+lenToPad)
	copy(padded, input)
	for i := len(input); i < len(padded); i++ {
		padded[i] = byte(lenToPad)
	}

	return padded
}
