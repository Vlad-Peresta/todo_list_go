package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Vlad-Peresta/todo_list_go/src/controllers"
)


// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	// route.GET("/todo", controllers.GetAllTodos)
	// route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	// route.DELETE("todo/:idTodo", controllers.DeleteTodo)

	// Run route whenever triggered
	route.Run()
}
