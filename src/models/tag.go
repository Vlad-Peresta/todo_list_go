package models

import (
	"errors"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
)

// Define Tag database table
type Tag struct {
	BaseModel
	Name  string  `json:"name"`
	Todos []*Todo `gorm:"many2many:todo_tags" json:"todos"`
}

func PatchUpdateTagByID[I int | string](tag *Tag, data *schemas.TagRequest, id I) error {
	if err := config.DB.First(&tag, "id = ?", id).Error; err != nil {
		return errors.New("tag record with provided ID was not found")
	}
	if err := config.DB.Model(&tag).Updates(map[string]any{
		"name": data.Name,
	}).Error; err != nil {
		return err
	}
	return nil
}
