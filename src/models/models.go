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

// Define Status database table
type Status struct {
	gorm.Model
	Name  string `json:"name"`
	Todos []Todo
}

// Define Tag database table
type Tag struct {
	gorm.Model
	Name  string  `json:"name"`
	Todos []*Todo `gorm:"many2many:todo_tags"`
}
