package utils

import (
	"crypto/sha256"
)

func GenerateHash(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return string(hash.Sum(nil))
}