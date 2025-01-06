package models

import (
	"errors"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
)

// GetAllRecords finds all records of the given model
func GetAllRecords[T interface{}](records *[]T) (err error) {
	if err := config.DB.Find(&records).Error; err != nil {
		return err
	}
	return nil
}

// GetRecordByID finds records of the given model by ID
func GetRecordByID[T interface{}, I uint | string](record *T, id I) (err error) {
	if err := config.DB.First(&record, "id = ?", id).Error; err != nil {
		return errors.New("Record with provided ID was not found.")
	}
	return nil
}

// CreateRecord creates record of the given model with provided data
func CreateRecord[T interface{}](record T) (err error) {
	if err := config.DB.Create(&record).Error; err != nil {
		return err
	}
	return nil
}

// PatchUpdateTodoByID updates todo record with provided data
func PatchUpdateTodoByID[I uint | string](todo *Todo, data schemas.TodoRequest, id I) (err error) {
	recordID := config.DB.First(&todo, "id = ?", id)
	if recordID.Error != nil {
		return errors.New("Todo record with provided ID was not found.")
	}
	if err := config.DB.Model(&todo).Updates(map[string]interface{}{
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

// DeleteRecordByID deletes record of the given model by ID
func DeleteRecordByID[T interface{}, I uint | string](record *T, id I) (err error) {
	recordID := config.DB.First(&record, "id = ?", id)
	if recordID.Error != nil {
		return errors.New("Record with provided ID was not found.")
	}

	if err := config.DB.Unscoped().Delete(&record).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByUsername finds record of the given model by `Username` field
func GetUserByUsername(record *User, username string) (err error) {
	if err := config.DB.First(&record, "username = ?", username).Error; err != nil {
		return errors.New("User with provided username was not found.")
	}
	return nil
}
