package models

import (
	"errors"
	"time"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
)

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

// PatchUpdateTodoByID updates todo record with provided data
func PatchUpdateTodoByID[I uint | string](todo *Todo, data schemas.TodoRequest, id I) (err error) {
	recordID := config.DB.First(&todo, "id = ?", id)
	if recordID.Error != nil {
		return errors.New("todo record with provided ID was not found")
	}
	if err := config.DB.Model(&todo).Updates(map[string]any {
		"name":        data.Name,
		"description": data.Description,
		"deadline":    data.Deadline,
		"active":      data.Active,
		"status_id":   data.StatusID,
		"user_id":     data.UserID,
	}).Error; err != nil {
		return err
	}
	return nil
}
