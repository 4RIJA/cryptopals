package cryptoutil

import "strings"

func Base64Encode(b []byte) string {
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
