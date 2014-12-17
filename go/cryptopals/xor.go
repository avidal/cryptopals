package cryptopals

import (
	"encoding/hex"
	"errors"
)

// Take two equal length byte arrays, xor them, return the result
func FixedXOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("a and b must be the same length")
	}

	sz := len(a)

	xor := make([]byte, sz)

	for i := range xor {
		xor[i] = a[i] ^ b[i]
	}

	return xor, nil
}

// Take two equal length hex strings, decode to byte arrays, and return the XOR
func FixedXORString(a, b string) (string, error) {
	a_hex, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}

	b_hex, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}

	xor, err := FixedXOR(a_hex, b_hex)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(xor), nil
}
