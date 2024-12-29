package controllers

import (
	"net/http"

	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// CreateTodo godoc
//
//	@Summary		Create Todo record
//	@Description	Create Todo record
//	@Tags			todos
//	@Produce		json
//	@Param			Request Body 	body		schemas.TodoRequest  	true	"Request Body"
//	@Success		200	{object}	schemas.TodoResponse
//	@Failure		400	{object}	error
//	@Router			/todos [POST]
//
// Create Todo record in database
func CreateTodo(context *gin.Context) {
	var data schemas.TodoRequest

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
	var response schemas.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Create HTTP response
	context.JSON(http.StatusCreated, response)
}

// GetAllTodos godoc
//
//	@Summary		Get all Todo records
//	@Description	Get all Todo records
//	@Tags			todos
//	@Produce		json
//	@Success		200	{object}	[]models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos [GET]
//
// Getting all Todo data
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	// Query to find all todos
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Error getting data"})
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

// GetTodo godoc
//
//	@Summary		Get Todo record by ID
//	@Description	Get Todo record by ID
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	schemas.TodoResponse
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [GET]
//
// Getting Todo record by ID
func GetTodo(context *gin.Context) {
	var todo models.Todo

	// Get Todo id from HTTP request parameter
	todoId := cast.ToUint(context.Param("id"))

	// Query to find todo
	err := db.First(&todo, todoId)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Record with provided ID does not exit"})
		return
	}

	// Matching result to todoResponse
	var response schemas.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating HTTP response
	context.JSON(http.StatusCreated, response)
}

// UpdateTodo godoc
//
//	@Summary		Update Todo record
//	@Description	Update Todo record
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Param			Request Body 	body		schemas.TodoRequest  	true	"Request Body"
//	@Success		200	{object}	schemas.TodoResponse
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [PUT]
//
// Update Todo record by ID
func UpdateTodo(context *gin.Context) {
	var data schemas.TodoRequest

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
	var response schemas.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating HTTP response
	context.JSON(http.StatusCreated, response)
}

// DeleteTodo godoc
//
//	@Summary		Delete Todo record
//	@Description	Delete Todo record
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	[]models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [DELETE]
//
// Delete Todo record by ID
func DeleteTodo(context *gin.Context) {
	// Initiate empty Todo model's record
	todo := models.Todo{}

	// Defining HTTP request parameter to get Todo id
	reqId := context.Param("id")
	todoId := cast.ToUint(reqId)

	// Delete Todo record by id from DB
	db.Where("id = ?", todoId).Unscoped().Delete(&todo)

	// Creating HTTP response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Record was deleted successfully",
		"data":    todoId,
	})
}
