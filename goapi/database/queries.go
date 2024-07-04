package database

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"

	"github.com/palSagnik/zephyr/models"

	"github.com/gofiber/fiber/v2"
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

func doesEmailExist(email string) (bool, error) {
	user := new(models.User)
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

func AddUser(c *fiber.Ctx, email string) (string, error) {
	
	// checking if the email already exists
	found, err := (doesEmailExist(email))
	if found {
		return "user already exists", errors.New("token already verified")
	}
	if err != nil {
		return err.Error(), err
	}

	// Get user from toverify table
	var toverify models.Verification
	result := DB.Select("username, email, password").Where("email = ?", email).First(&toverify)
	if result.Error != nil {
		return "token expired, please register again", result.Error
	}

	// Create user
	result = DB.Create(&models.User{Username: toverify.Username, Email: toverify.Email, Password: toverify.Password})
	if result.Error != nil {
		return "error in creating user, please contact admin", result.Error
	}

	// After creating user delete from toverify table
	_ = DB.Delete(&models.Verification{}, "email = ?", email)

	return "", nil
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

func DeleteRunning(c *fiber.Ctx, userid int) error {
	result := DB.Delete(&models.RunningInstance{}, "userid = ?", userid)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
