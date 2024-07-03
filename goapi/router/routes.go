package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/zephyr/handler"
	"github.com/palSagnik/zephyr/middleware"
)

func SetUpRoutes (app *fiber.App) {
	app.Get("/alive", handler.Alive)

	auth := app.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Post("/login", handler.Login)

	api := app.Group("api", middleware.VerifyToken())
	api.Post("/configurations", handler.ConfigList)
	// api.Get("/instance", handler.StartInstance)

}