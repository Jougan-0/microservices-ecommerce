package main

import (
	"user-service/db"
	"user-service/routes"
	"user-service/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InitLogger()
	db.InitDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	utils.Logger.Info("Starting User Service on port 3001")
	app.Listen(":3001")
}
