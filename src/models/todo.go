package models

import (
	"time"

	"gorm.io/gorm"
)

// Define Todo database table
type Todo struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Active      bool      `gorm:"default:true" json:"active"`
	Tags        []*Tag    `gorm:"many2many:todo_tags"`
	StatusID    uint
}

// func GetAllTodos(t *[]Todo) (err error) {

// }
