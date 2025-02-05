package routes

import (
	"user-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.LoginUser)
}
