package routes

import (
	"github.com/gofiber/fiber/v2"

	"crud/handlers"
)

func SetupRoutes(app *fiber.App) {
	items := app.Group("/items")
	items.Get("/", handlers.GetAllItems)
	items.Post("/", handlers.InsertNewItems)
}