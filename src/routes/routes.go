package routes

import (
	"fmt"

	"github.com/Vlad-Peresta/todo_list_go/src/controllers"
	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo/:id", controllers.UpdateTodo)
	// route.DELETE("todo/:idTodo", controllers.DeleteTodo)

	// Run route whenever triggered
	if err := route.Run(); err != nil {
		fmt.Printf("Failed to start the server: %s", err)
		return
	}
}
