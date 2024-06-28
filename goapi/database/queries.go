package database

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/zephyr/models"
)

// Random string generator for instance passwords
func GenerateRandom() string {
	buffer := make([]byte, 128)
	rand.Read(buffer)

	return hex.EncodeToString(buffer)
}

// miscellanous queries
func VerifyCreds(c *fiber.Ctx) error {

	return nil
}

func DoesEmailExist(email string) bool {
	var uid int
	if err := DB.QueryRow(`SELECT userid FROM users WHERE email = $1`, email).Scan(&uid); err == nil {
		return true
	}
	return false
}

// user related queries
func AddUser(c *fiber.Ctx, email string) (string, error) {
	ctx, cancel := context.WithTimeout(c.Context(), 15*time.Second)
	defer cancel()

	userData := new(models.User)

	// checking if the email already exists
	if DoesEmailExist(email) {
		return "user already exists", errors.New("user exists")
	}

	if _, err := DB.QueryContext(ctx, `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, userData.Username, userData.Email, userData.Password); err != nil {
		log.Println(err.Error())
		return "user could not be added", err
	}

	return "user added successfully", nil
}

func DoesUserExist(c *fiber.Ctx, userid int) bool {
	var username int
	ctx, cancel := context.WithTimeout(c.Context(), 15*time.Second)
	defer cancel()

	if err := DB.QueryRowContext(ctx, `SELECT username FROM users WHERE userid = $1`, userid).Scan(&username); err != nil {
		return false
	}
	return true
}


// Instance related queries
func CanStartInstance(c *fiber.Ctx, userid int) bool {
	var runid int	
	ctx, cancel := context.WithTimeout(c.Context(), 15*time.Second)
	defer cancel()

	if err := DB.QueryRowContext(ctx, `SELECT runid FROM running WHERE userid = $1`, userid).Scan(&runid); err != nil {
		return false
	}

	if _, err := DB.QueryContext(ctx, `INSERT INTO running(userid) VALUES ($1)`, userid); err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func DeleteRunningInstance(c *fiber.Ctx, userid int) error {
	ctx, cancel := context.WithTimeout(c.Context(), 15*time.Second)
	defer cancel()

	if _, err := DB.QueryContext(ctx, `DELETE FROM running WHERE userid = $1`, userid); err != nil {
		return err
	}

	return nil
}


