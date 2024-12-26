package controllers

import (
	"net/http"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// Todo struct for the request HTTP body
type todoRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// Todo struct for the HTTP response
type todoResponse struct {
	todoRequest
	ID uint `json:"ID"`
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

func UpdateTodo(context *gin.Context) {
	var data todoRequest

	// Defining HTTP request parameter to get Todo id
	reqId := context.Param("id")
	todoId := cast.ToUint(reqId)

	// Binding HTTP request body to the todoRequest struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate empty Todo model's record
	todo := models.Todo{}

	// Get first Todo record by id from DB
	todoById := db.Where("id = ?", todoId).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	// Matching todoRequest with models.Todo
	todo.Name = data.Name
	todo.Description = data.Description

	// Update existing Todo record
	result := db.Save(&todo)
	if err := result.Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Matching result to todoResponse
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

  // Creating HTTP response
	context.JSON(http.StatusCreated, response)
}
