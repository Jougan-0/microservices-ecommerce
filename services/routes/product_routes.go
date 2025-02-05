package routes

import (
	"microservices/handlers"
	"microservices/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCatalogRoutes(app *fiber.App) {
	userGroup := app.Group("/api", middleware.JWTAuthMiddleware)

	userGroup.Post("/product", handlers.CreateProduct)
	userGroup.Put("/product/:id", handlers.UpdateProduct)
	userGroup.Delete("/product/:id", handlers.DeleteProduct)
	app.Get("/products", handlers.GetProducts)
	app.Get("/product/:id", handlers.GetProductByID)
}
