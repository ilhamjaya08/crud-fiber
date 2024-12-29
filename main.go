package main

import (
	"crud/route"
	"crud/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "Crud Application",
	})

	app.Use(logger.New())

	route.SetupRoutes(app)

	app.Listen(":3000")
}