package database

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/zephyr/models"
	"gorm.io/gorm"
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

func doesEmailExist(user *models.User, email string) (bool, error) {
	
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			// Record not found
			return false, nil
		}
		// Other error
		return false, result.Error
	}
	
	return true, nil
}

// user related queries
func DeleteUser(c *fiber.Ctx) error {
	
	email := c.Params("email")
	result := DB.Delete(&email)
	if result.Error != nil {
		return result.Error
	}
	
	return nil
}

func AddUser(c *fiber.Ctx) error {

	userData := new(models.User)
	userData.Email = c.Params("email")
	userData.Username = c.Params("username")
	userData.Password = c.Params("password")

	// checking if the email already exists
	found, err := (doesEmailExist(userData, userData.Email))
	if found {
		return errors.New("email already exists")
	} 
	if err != nil {
		return err
	}

	result := DB.Create(&userData)
	if result.Error != nil {
		return result.Error
	}

	// TODO: After creating user delete from toverify table

	return nil
	
}

func DoesUserExist(c *fiber.Ctx) bool {
	userid := c.Params("userid")

	user := new(models.User)
	result := DB.Find(&user, userid)

	if result.RowsAffected == 0 {
		return false
	}
	return true
}


// Instance related queries
func GetAllRunningInstances(c *fiber.Ctx) error {
	var instances []models.RunningInstance

	if result := DB.Find(&instances); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status":"failure", "message":"no instances found"})
	}

	return c.Status(fiber.StatusOK).JSON(instances)
}

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




