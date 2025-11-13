package tools

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHash(hash string, raw string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])

	return hashStr
}

func VerifyHashSHA256(input string, expectedHash string) bool {
	inputHash := HashSHA256(input)

	return inputHash == expectedHash
}

func GenerateSecureTokenHex(n int) (string, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
