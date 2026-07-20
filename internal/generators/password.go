package generators

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

const (
	letterCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numberCharset = "0123456789"
	symbolCharset = "!@#$%^&*()-_=+"
)

// GeneratePassword returns a random password of the given length using a
// cryptographically secure random source.
func GeneratePassword(length int, noLetters bool, noNumber bool, noSymbol bool) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be positive, got %d", length)
	}

	charSet := ""

	if !noLetters{
		charSet = letterCharset
	}
	if !noNumber{
		charSet += numberCharset
	}
	if !noSymbol{
		charSet += symbolCharset
	}

	if len(charSet) == 0 {
		return "", fmt.Errorf("no characters available: cannot exclude all character sets")
	}

	result := make([]byte, length)
	charsetSize := big.NewInt(int64(len(charSet)))

	for i := range result {
		idx, err := rand.Int(rand.Reader, charsetSize)
		
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		result[i] = charSet[idx.Int64()]
	}

	return string(result), nil
}