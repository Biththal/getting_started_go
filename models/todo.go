package models

import (
	"todo/config"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func MigrateTodo() {
	config.DB.AutoMigrate(&Todo{})
}
