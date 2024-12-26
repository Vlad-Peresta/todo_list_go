package controllers

import (
	"net/http"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// Todo struct for the request HTTP body
type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"desctiption"`
}

// Todo struct for the HTTP response
type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

// Create Todo record in database
func CreateTodo(context *gin.Context) {
	var data todoRequest

	// Binding JSON request body to todoRequest struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching Todo models struct to todoRequest struct
	todo := models.Todo{Name: data.Name, Description: data.Description}

	// Query to database
	result := db.Create(&todo)
	if err := result.Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Matching result to create HTTP Response
	// response := todoResponse{ID: todo.ID, Name: todo.Name, Description: todo.Description}
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Create HTTP response
	context.JSON(http.StatusCreated, response)
}

// Getting all Todo data
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	// Query to find all todos
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating HTTP response
	context.JSON(
		http.StatusOK, gin.H{
			"status":  "200",
			"message": "Success",
			"data":    todos,
		})
}
