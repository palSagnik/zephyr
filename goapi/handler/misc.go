package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Alive(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Server is up!"})
}