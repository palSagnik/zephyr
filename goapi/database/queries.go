package database

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"

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
func ValidateCreds(c *fiber.Ctx, creds *models.Credentials) error {
	var user models.User

	result := DB.Where("email = ? AND password = ?", creds.Email, creds.Password).First(&user)
	if result.Error != nil {
		return result.Error
	}
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
func AddToVerify(c *fiber.Ctx, user *models.User) error {
	email := user.Email
	result := DB.Delete(&models.Verification{}, "email =  ?", email)
	if result.Error != nil {
		return err
	}

	result = DB.Create(&models.Verification{Username: user.Username, Email: email, Password: user.Password})
	if result.Error != nil {
		return err
	}
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	email := c.Params("email")
	result := DB.Delete(&models.User{}, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddUser(c *fiber.Ctx, user *models.User) error {
	
	// checking if the email already exists
	found, err := (doesEmailExist(user, user.Email))
	if found {
		return errors.New("email already exists")
	}
	if err != nil {
		return err
	}

	result := DB.Create(&models.User{Username: user.Username, Email: user.Email, Password: user.Password})
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

	return result.RowsAffected > 0
}

// Instance related queries
func GetInstances(c *fiber.Ctx) error {
	var instances []models.RunningInstance

	if result := DB.Find(&instances); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "failure", "message": "no instances found"})
	}

	return c.Status(fiber.StatusOK).JSON(instances)
}

func CanStartInstance(c *fiber.Ctx, userid int) bool {
	result := DB.Find(&models.RunningInstance{}, "userid = ?", userid)
	if result.Error == nil {
		log.Println(err)
		return false
	}

	result = DB.Create(&models.RunningInstance{UserID: userid})
	if result.Error != nil {
		log.Println(err)
		return false
	}

	return true
}

func DeleteRunning (c *fiber.Ctx, userid int) error {
	result := DB.Delete(&models.RunningInstance{}, "userid = ?", userid)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
