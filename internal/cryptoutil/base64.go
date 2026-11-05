package cryptoutil

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Base64Encode(b []byte) string {
	var sol strings.Builder
	full := len(b) - len(b)%3

	for i := 0; i < full; i += 3 {
		sol.WriteByte(binToBase64(b[i] >> 2))
		sol.WriteByte(binToBase64((b[i]&0x03)<<4 | b[i+1]>>4))
		sol.WriteByte(binToBase64((b[i+1]&0x0F)<<2 | b[i+2]>>6))
		sol.WriteByte(binToBase64(b[i+2] & 0x3F))
	}

	switch len(b) % 3 {
	case 1:
		x := b[full]
		sol.WriteByte(binToBase64(x >> 2))
		sol.WriteByte(binToBase64((x & 0x03) << 4))
		sol.WriteString("==")
	case 2:
		x, y := b[full], b[full+1]
		sol.WriteByte(binToBase64(x >> 2))
		sol.WriteByte(binToBase64((x&0x03)<<4 | y>>4))
		sol.WriteByte(binToBase64((y & 0x0F) << 2))
		sol.WriteByte('=')
	}

	return sol.String()
}

func binToBase64(b byte) byte {
	switch {
	case b < 26:
		return 'A' + b
	case b < 52:
		return 'a' + (b - 26)
	case b < 62:
		return '0' + (b - 52)
	case b == 62:
		return '+'
	default:
		return '/'
	}
}

func Base64Decode(toDecode []byte) ([]byte, error) {
	result := make([]byte, base64.StdEncoding.DecodedLen(len(toDecode)))
	n, err := base64.StdEncoding.Decode(result, toDecode)
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}
	return result[:n], nil
}
