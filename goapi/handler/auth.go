package handler

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/zephyr/database"
	"github.com/palSagnik/zephyr/middleware"
	"github.com/palSagnik/zephyr/models"
	"github.com/palSagnik/zephyr/utils"
)

func Signup(c *fiber.Ctx) error{
	
	signupForm := new(models.User)

	// assign form values
	signupForm.Email = c.FormValue("email")
	signupForm.Username = c.FormValue("username")
	signupForm.Password = c.FormValue("password")
	signupForm.ConfirmPass = c.FormValue("confirm")

	// handle error when any field is empty
	if signupForm.Email == "" || signupForm.Username == "" || signupForm.Password == "" || signupForm.ConfirmPass == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":"all fields must be filled"})
	}

	// format form values
	signupForm.Email = strings.TrimSpace(signupForm.Email)
	signupForm.Username = strings.TrimSpace(signupForm.Username)
	signupForm.Password = strings.TrimSpace(signupForm.Password)
	signupForm.ConfirmPass = strings.TrimSpace(signupForm.ConfirmPass)

	signupForm.Email = strings.ToLower(signupForm.Email)

	if isOk, statusMsg := utils.VerifySignUpInput(signupForm); !isOk {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":statusMsg})
	}

	// send verification mail
	if err := utils.SendVerificationMail(signupForm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"failed", "message":"error in sending verification email"})
	}

	// store hash of the password
	signupForm.Password = utils.GenerateHash(signupForm.Password)

	if err := database.AddToVerify(c, signupForm); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"failed", "message":"please contact admin"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"success", "message":"check your mail for verification"})
}

func Login(c *fiber.Ctx) error {
	creds := new(models.Credentials)
	user := new(models.User)

	creds.Email = c.FormValue("email")
	creds.Password = c.FormValue("password")

	//basic check for empty fields
	if creds.Email == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":"email and password required"})
	}

	// simple sanitisation and verification
	creds.Email = strings.TrimSpace(creds.Email)
	creds.Password = strings.TrimSpace(creds.Password)
	creds.Email = strings.ToLower(creds.Email)

	if isOk, statusMsg := utils.VerifyLoginInput(creds); !isOk {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"failed", "message":statusMsg})
	}

	// hashing the password
	creds.Password = utils.GenerateHash(creds.Password)

	// checking email and password
	if err := database.ValidateCreds(c, creds); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "failure", "message": "invalid credentials"})
	}

	// TODO: token generation and assigning
	token, err := middleware.GenerateToken(user)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "message": "error in token generation. contact admin"})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt-token"
	cookie.Value = token
	cookie.SameSite = fiber.CookieSameSiteStrictMode
	cookie.Expires = time.Now().Add(72 * time.Hour)

	c.Cookie(cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "login successful"})
}