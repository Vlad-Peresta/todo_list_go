package models

import (
	"time"

	"gorm.io/gorm"
)

// Define Status database table
type Status struct {
	gorm.Model
	Name  string
	Todos []Todo
}

// Define Tag database table
type Tag struct {
	gorm.Model
	Name  string
	Todos []*Todo `gorm:"many2many:todo_tags"`
}

// Define Todo database table
type Todo struct {
	gorm.Model
	Name        string
	Description string
	Deadline    time.Time
	Active      bool   `gorm:"default:true"`
	Tags        []*Tag `gorm:"many2many:todo_tags"`
	StatusID    uint
}
