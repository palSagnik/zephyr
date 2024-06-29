package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
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

	// TODO: ADD TO VERIFY TABLE
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"success", "message":"check your mail for verification"})

}