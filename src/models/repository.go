package models

import (
	"errors"
	"time"

	"gorm.io/gorm"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
)

type Model interface {
	User | Todo | Tag | Status
}

type ID interface {
	uint | string
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// GetAllRecords finds all records of the given model
func GetAllRecords[M Model](records *[]M) (err error) {
	if err := config.DB.Find(&records).Error; err != nil {
		return err
	}
	return nil
}

// GetRecordByID finds records of the given model by ID
func GetRecordByID[M Model, I ID](record *M, id I) (err error) {
	if err := config.DB.First(&record, "id = ?", id).Error; err != nil {
		return errors.New("record with provided ID was not found")
	}
	return nil
}

// CreateRecord creates record of the given model with provided data
func CreateRecord[M Model](record *M) (err error) {
	if err := config.DB.Create(&record).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRecordByID deletes record of the given model by ID
func DeleteRecordByID[M Model, I ID](record *M, id I) (err error) {
	recordID := config.DB.First(&record, "id = ?", id)
	if recordID.Error != nil {
		return errors.New("record with provided ID was not found")
	}

	if err := config.DB.Unscoped().Delete(&record).Error; err != nil {
		return err
	}
	return nil
}
