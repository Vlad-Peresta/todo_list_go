package models

import (
	"errors"
	"math"
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

type Pagination struct {
	Limit        int    `json:"limit,omitempty" form:"limit"`
	Page         int    `json:"page,omitempty" form:"page"`
	Sort         string `json:"sort,omitempty" form:"sort"`
	TotalRecords int64  `json:"total_records"`
	TotalPages   int    `json:"total_pages"`
	Data         any    `json:"data"`
}

func (p *Pagination) GetPaginatedAndSortedRecords() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.getOffset()).Limit(p.getLimit()).Order(p.getSort())
	}
}

func (p *Pagination) getPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) getSort() string {
	if p.Sort == "" {
		return "id DESC"
	}
	return p.Sort
}

func (p *Pagination) getOffset() int {
	return (p.getPage() - 1) * p.Limit
}

func (p *Pagination) getLimit() int {
	if p.Limit == 0 {
		return 100
	}
	return p.Limit
}

// GetAllRecords finds all records of the given model
func GetAllRecords[M Model](records *[]M, pagination *Pagination) (err error) {
	var totalRecords int64
	if err := config.DB.Scopes(pagination.GetPaginatedAndSortedRecords()).Find(&records).Error; err != nil {
		return err
	}
	config.DB.Model(&records).Count(&totalRecords)
	pagination.TotalRecords = totalRecords
	pagination.TotalPages = int(math.Ceil(float64(totalRecords) / float64(pagination.getLimit())))
	pagination.Data = records
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
