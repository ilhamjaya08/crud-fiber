package main

import (
	"crud/database"
	"crud/migrations"
	"crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	migrations.Migrate()

	app := fiber.New(fiber.Config{
		AppName: "Crud Application",
	})

	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}