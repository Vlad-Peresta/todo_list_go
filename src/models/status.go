package models

// Define Status database table
type Status struct {
	BaseModel
	Name  string `json:"name"`
	Todos []Todo
}
