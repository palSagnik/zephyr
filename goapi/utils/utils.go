package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"net/mail"

	"github.com/palSagnik/zephyr/config"
	"github.com/palSagnik/zephyr/database"
	"github.com/palSagnik/zephyr/models"
)

func UpdateKey(key string) error {
	secret := database.GenerateRandom()
	err := os.Setenv(key, secret)
	if err != nil {
		return err
	}
	return nil
}

func GenerateHash(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func VerifyLoginInput(creds *models.Credentials) (bool, string) {
	emailLength := len(creds.Email)
	if emailLength > config.MAIL_LEN {
		return false, fmt.Sprintf("Email should not exceed %d characters", config.MAIL_LEN)
	}

	passLength := len(creds.Password)
	if passLength > config.PASS_LEN || passLength < 8 {
		return false, fmt.Sprintf("Password should be of length 8-%d characters", config.PASS_LEN)
	}

	if _, err := mail.ParseAddress(creds.Email); err != nil {
		return false, "Not a valid email address"
	}
	return true, ""
}

func VerifySignUpInput(signupForm *models.User) (bool, string) {

	passLength := len(signupForm.Password)
	if passLength > config.PASS_LEN || passLength < 8 {
		return false, fmt.Sprintf("Password should be of length 8-%d characters", config.PASS_LEN)
	}
	
	if signupForm.ConfirmPass != signupForm.Password {
		return false, "Passwords don't match"
	}

	usernameLength := len(signupForm.Username)
	if usernameLength > config.USERNAME_LEN || usernameLength < 4 {
		return false, fmt.Sprintf("Username should be of length 4-%d characters", config.USERNAME_LEN)
	}

	emailLength := len(signupForm.Email)
	if emailLength > config.MAIL_LEN {
		return false, fmt.Sprintf("Email should not exceed %d characters", config.MAIL_LEN)
	}

	// TODO: Email address parsing verify
	if _, err := mail.ParseAddress(signupForm.Email); err != nil {
		return false, "Not a valid email address"
	}

	// TODO: verify password regex
	
	return true, ""
	
}

func SendVerificationMail(signsignupForm *models.User) error {
	log.Println("verification email sent")
	return nil
}