package main

import (
	"user-service/db"
	"user-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":3001")
}
