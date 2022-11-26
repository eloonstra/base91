package base91

import (
	"bytes"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

// Encode encodes a byte slice into a base91 string
func Encode(data []byte) string {
	var (
		b   int
		n   uint8
		out strings.Builder
	)

	for _, c := range data {
		b |= int(c) << n
		n += 8
		if n > 13 {
			v := b & 8191
			if v > 88 {
				b >>= 13
				n -= 13
			} else {
				v = b & 16383
				b >>= 14
				n -= 14
			}
			out.WriteString(string(alphabet[v%91]) + string(alphabet[v/91]))
		}
	}

	if n > 0 {
		out.WriteByte(alphabet[b%91])
		if n > 7 || b > 90 {
			out.WriteByte(alphabet[b/91])
		}
	}

	return out.String()
}

// Decode decodes a base91 string into a byte slice
func Decode(data string) []byte {
	var (
		v           = -1
		b           int
		n           uint8
		out         bytes.Buffer
		decodeTable = genDecodeTable()
	)

	for _, l := range data {
		if _, ok := decodeTable[string(l)]; !ok {
			continue
		}

		c := decodeTable[string(l)]
		if v < 0 {
			v = int(c)
			continue
		}

		v += int(c) * 91
		b |= v << n
		if v&8191 > 88 {
			n += 13
		} else {
			n += 14
		}

		for {
			out.WriteByte(byte(b & 255))
			b >>= 8
			n -= 8
			if n <= 7 {
				break
			}
		}

		v = -1
	}

	if v+1 > 0 {
		out.WriteByte(byte(b|v<<n) & 255)
	}

	return out.Bytes()
}

func genDecodeTable() map[string]uint8 {
	decodeTable := make(map[string]uint8)
	for i, c := range alphabet {
		decodeTable[string(c)] = uint8(i)
	}
	return decodeTable
}
