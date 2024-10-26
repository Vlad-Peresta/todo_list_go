package models

import (
	"gorm.io/gorm"
)

// Define Todo table for database communications
type Todo struct {
	gorm.Model
	Name        string
	Description string
}
