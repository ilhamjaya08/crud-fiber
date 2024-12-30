package handlers

import (
	"crud/database"
	"crud/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	db := database.DB
	var items []models.Item

	db.Find(&items)
	return c.JSON(fiber.Map{
		"status": "success",
		"data": items,
	})
}

func InsertNewItems(c *fiber.Ctx) error {
	db:= database.DB
	item := new(models.Item)

	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"message": "Invalid request",
		})
	}

	db.Create(&item)
	return c.JSON(fiber.Map{
		"status": "success",
		"data": item,
	})
}