package migrations

import (
	"crud/models"
	"log"
	"crud/database"
)

func Migrate() {
	err := database.DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database. \n", err)
	}
	log.Println("Migrated Succesfully")
}