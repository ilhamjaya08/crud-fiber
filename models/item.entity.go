package models


import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}