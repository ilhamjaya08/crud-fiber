package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/crud?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database, error: " + err.Error())
	}
	DB = db

	fmt.Println("Connected to database")
}