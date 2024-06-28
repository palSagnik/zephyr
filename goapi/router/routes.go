package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/zephyr/handler"
)

func SetUpRoutes (app *fiber.App) {
	app.Get("/alive", handler.Alive)

	auth := app.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Get("/login", handler.Login)

	api := app.Group("api", middleware.VerifyToken)
	api.Post("/configuration", handler.ConfigList)
	api.Get("/instance", handler.StartInstance)

}