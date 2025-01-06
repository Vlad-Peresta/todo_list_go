package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Define Todo database table
type Todo struct {
	BaseModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Active      bool      `gorm:"default:true" json:"active"`
	Tags        []*Tag    `gorm:"many2many:todo_tags" json:"tags"`
	StatusID    uint      `json:"status_id"`
	UserID      uint      `json:"user_id"`
}

// Define Status database table
type Status struct {
	BaseModel
	Name  string `json:"name"`
	Todos []Todo
}

// Define Tag database table
type Tag struct {
	BaseModel
	Name  string  `json:"name"`
	Todos []*Todo `gorm:"many2many:todo_tags"`
}

// Define User database table
type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Todos    []Todo
}
