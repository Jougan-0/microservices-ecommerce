package main

import (
	"microservices/db"
	"microservices/routes"
	"microservices/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.LoadEnv()

	utils.InitLogger()
	db.InitDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                           // Allows requests from any domain
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS", // Allowed HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.SetupUserRoutes(app)
	routes.SetupCatalogRoutes(app)
	routes.SetupCartRoutes(app)

	utils.Logger.Info("Starting User Service on port 3001")
	app.Listen(":3001")
}
