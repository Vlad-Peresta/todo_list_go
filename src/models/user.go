package models

import (
	"errors"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"gorm.io/gorm"
)

// Define User database table
type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Todos    []Todo
}

// GetUserByUsername finds record of the given model by `Username` field
func GetUserByUsername(record *User, username string) (err error) {
	if err := config.DB.First(&record, "username = ?", username).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user with provided username was not found")
	}
	return nil
}
