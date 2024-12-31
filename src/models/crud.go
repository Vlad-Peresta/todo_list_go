package models

import (
	"errors"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
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

// PatchUpdateRecordByID updates record of the given model with provided data
func PatchUpdateRecordByID[T interface{}, D interface{}, I uint | string](record *T, data D, id I) (err error) {
	recordID := config.DB.First(&record, "id = ?", id)
	if recordID.Error != nil {
		return errors.New("Record with provided ID was not found.")
	}
	if err := config.DB.Model(&record).Updates(data).Error; err != nil {
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
