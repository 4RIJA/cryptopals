package challenge1

import (
	"errors"
	"strings"
)

func hexToBase64(s string) (string, error) {
	bin, err := hexDecode(s)
	if err != nil {
		return "", err
	}
	return base64Encode(bin), nil
}

func hexDecode(s string) ([]byte, error) {
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

func base64Encode(b []byte) string {
	var sol strings.Builder
	full := len(b) - len(b)%3

	for i := 0; i < full; i += 3 {
		sol.WriteByte(base64Value(b[i] >> 2))
		sol.WriteByte(base64Value((b[i]&0x03)<<4 | b[i+1]>>4))
		sol.WriteByte(base64Value((b[i+1]&0x0F)<<2 | b[i+2]>>6))
		sol.WriteByte(base64Value(b[i+2] & 0x3F))
	}

	switch len(b) % 3 {
	case 1:
		x := b[full]
		sol.WriteByte(base64Value(x >> 2))
		sol.WriteByte(base64Value((x & 0x03) << 4))
		sol.WriteString("==")
	case 2:
		x, y := b[full], b[full+1]
		sol.WriteByte(base64Value(x >> 2))
		sol.WriteByte(base64Value((x&0x03)<<4 | y>>4))
		sol.WriteByte(base64Value((y & 0x0F) << 2))
		sol.WriteByte('=')
	}

	return sol.String()
}

func base64Value(b byte) byte {
	switch {
	case b < 0x1A:
		return 'A' + b
	case b < 0x34:
		return 'a' + (b - 0x1A)
	case b < 0x3E:
		return '0' + (b - 0x34)
	case b == 0x3E:
		return '+'
	default:
		return '/'
	}
}
