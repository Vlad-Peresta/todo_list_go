package models

// Define Tag database table
type Tag struct {
	BaseModel
	Name  string  `json:"name"`
	Todos []*Todo `gorm:"many2many:todo_tags"`
}
