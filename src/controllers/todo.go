package controllers

import (
	"net/http"

	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateTodo godoc
//
//	@Summary		Create Todo record
//	@Description	Create Todo record
//	@Tags			todos
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			Request Body 	body		schemas.TodoRequest  	true	"Request Body"
//	@Success		200	{object}	schemas.TodoResponse
//	@Failure		400	{object}	error
//	@Router			/todos [POST]
//
// CreateTodo creates Todo record in the database
func CreateTodo(context *gin.Context) {
	var todo models.Todo

	// Binding JSON request body to todo struct
	if err := context.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query to database
	err := models.CreateRecord(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Matching result to create HTTP Response
	response := schemas.TodoResponse{}
	if err := context.ShouldBindBodyWith(&response, binding.JSON); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.ID = todo.ID

	// Create HTTP response
	context.JSON(http.StatusCreated, response)
}

// GetAllTodos godoc
//
//	@Summary		Get all Todo records
//	@Description	Get all Todo records
//	@Tags			todos
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
// @Param        sort    query     string  false  "Sorting parameter"  default(id DESC)
// @Param        limit    query     string  false  "Records per page" default(100)
// @Param        page    query     string  false  "Current page" default(1)
//
//	@Produce		json
//	@Success		200	{object}	[]models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos [GET]
//
// GetAllTodos finds all Todo records
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo
	var pagination models.Pagination

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find all todo's records
	err := models.GetAllRecords(&todos, &pagination)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [GET]
//
// GetTodo finds Todo record by ID
func GetTodo(context *gin.Context) {
	var todo models.Todo

	// Finding todo record by id
	err := models.GetRecordByID(&todo, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creating HTTP response
	context.JSON(http.StatusCreated, todo)
}

// UpdateTodo godoc
//
//	@Summary		Update Todo record
//	@Description	Update Todo record
//	@Tags			todos
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Param			Request Body 	body		schemas.TodoRequest  	true	"Request Body"
//	@Success		200	{object}	models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [PUT]
//
// UpdateTodo updates Todo record by ID
func UpdateTodo(context *gin.Context) {
	var data schemas.TodoRequest

	// Binding HTTP request body to the todoRequest struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate empty Todo model's record
	todo := models.Todo{}

	// // Updating Todo record by id
	err := models.PatchUpdateTodoByID(&todo, data, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creating HTTP response
	context.JSON(http.StatusCreated, todo)
}

// DeleteTodo godoc
//
//	@Summary		Delete Todo record
//	@Description	Delete Todo record
//	@Tags			todos
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	[]models.Todo
//	@Failure		400	{object}	error
//	@Router			/todos/{id} [DELETE]
//
// DeleteTodo deletes Todo record by ID
func DeleteTodo(context *gin.Context) {
	// Initiate empty Todo model's record
	todo := models.Todo{}
	id := context.Param("id")

	// Delete Todo record by id from DB
	err := models.DeleteRecordByID(&todo, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Creating HTTP response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "record was deleted successfully",
		"data":    id,
	})
}
