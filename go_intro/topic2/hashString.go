package topic2

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(text string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(text))

	if err != nil {
		return "", err
	}

	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes), nil
}
