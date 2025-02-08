package routes

import (
	"microservices/handlers"
	"microservices/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCartRoutes(app *fiber.App) {
	cart := app.Group("/cart", middleware.JWTAuthMiddleware)

	cart.Post("/add", handlers.AddToCart)
	cart.Get("/", handlers.GetCart)
	cart.Delete("/remove/:id", handlers.RemoveCartItem)
	cart.Delete("/clear", handlers.ClearCart)
}
