package cryptoutil

import (
	"errors"
	"strings"
)

func HexDecode(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, errors.New("hex string has odd length")
	}

	decoded := make([]byte, len(s)/2)

	for i := 0; i < len(s); i += 2 {
		hi, err := hexValue(s[i])
		if err != nil {
			return nil, err
		}
		lo, err := hexValue(s[i+1])
		if err != nil {
			return nil, err
		}
		decoded[i/2] = hi<<4 | lo
	}

	return decoded, nil
}

func hexValue(c byte) (byte, error) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', nil
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, nil
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10, nil
	default:
		return 0, errors.New("invalid hexadecimal character")
	}
}

func HexEncode(b []byte) string {
	var sol strings.Builder
	for _, v := range b {
		v1, _ := valueHex(v >> 4)
		sol.WriteByte(v1)
		v2, _ := valueHex(v & 0b00001111)
		sol.WriteByte(v2)
	}
	return sol.String()
}

func valueHex(b byte) (byte, error) {
	switch {
	case b < 0x0A:
		return '0' + b, nil
	case b < 0x10:
		return 'a' + (b - 0x0A), nil
	default:
		return 0, errors.New("not a valid value to convert to hex")
	}
}
