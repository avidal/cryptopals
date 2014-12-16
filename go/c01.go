package cryptopals

// Take a string of hex digits, convert to base64
import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))

	return encoded, nil

}
