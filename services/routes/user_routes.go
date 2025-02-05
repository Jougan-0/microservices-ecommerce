package routes

import (
	"microservices/handlers"
	"microservices/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	// Public Routes
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.LoginUser)

	// Protected Routes (Require JWT authentication)
	userGroup := app.Group("/user", middleware.JWTAuthMiddleware)
	userGroup.Get("/profile", handlers.GetUserProfile)
	userGroup.Put("/update-profile", handlers.UpdateUserProfile)
	userGroup.Delete("/delete-account", handlers.DeleteUserAccount)
}
