package controllers

import (
	"net/http"

	"github.com/Vlad-Peresta/todo_list_go/src/config"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// Todo struct to the request body
type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"desctiption"`
}

// Todo struct for response
type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

// Create todo data to database by run this function
func CreateTodo(context *gin.Context) {
	var data todoRequest

	// Binding json request body  with with todoRequest struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching todo models struct with todo request struct
	todo := models.Todo{Name: data.Name, Description: data.Description}

	// Query to database
	result := db.Create(&todo)
	if err := result.Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Matching result to create Response
	response := todoResponse{ID: todo.ID, Name: todo.Name, Description: todo.Description}

	// Creating Http response
	context.JSON(http.StatusCreated, response)
}
