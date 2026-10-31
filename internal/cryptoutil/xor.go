package cryptoutil

import "errors"

func FixedXor(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("not the same length")
	}

	sol := make([]byte, len(b))
	for i := range a {
		sol[i] = xor(a[i], b[i])
	}
	return sol, nil
}

func xor(a byte, b byte) byte {
	return a ^ b
}
