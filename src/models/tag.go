package models

import (
	"gorm.io/gorm"
)

// Define Tag database table
type Tag struct {
	gorm.Model
	Name  string  `json:"name"`
	Todos []*Todo `gorm:"many2many:todo_tags"`
}
