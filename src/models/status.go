package models

import (
	"gorm.io/gorm"
)

// Define Status database table
type Status struct {
	gorm.Model
	Name  string `json:"name"`
	Todos []Todo
}
