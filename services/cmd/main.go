package main

import (
	"microservices/db"
	"microservices/routes"
	"microservices/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()

	utils.InitLogger()
	db.InitDB()

	app := fiber.New()
	routes.SetupUserRoutes(app)
	routes.SetupCatalogRoutes(app)

	utils.Logger.Info("Starting User Service on port 3001")
	app.Listen(":3001")
}
