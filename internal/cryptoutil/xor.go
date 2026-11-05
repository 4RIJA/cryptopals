package cryptoutil

import "errors"

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
