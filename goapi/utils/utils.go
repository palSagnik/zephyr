package utils

import (
	"crypto/sha256"

	"github.com/palSagnik/zephyr/models"
)

func GenerateHash(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return string(hash.Sum(nil))
}

func VerifySignUpInput(signupForm *models.User) (error, int) {

	
}

func SendVerificationMail(signsignupForm *models.User) error {
	return nil
}