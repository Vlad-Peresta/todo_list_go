package routes

import (
	"fmt"

	docs "github.com/Vlad-Peresta/todo_list_go/docs"
	"github.com/Vlad-Peresta/todo_list_go/src/controllers"
	"github.com/Vlad-Peresta/todo_list_go/src/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := route.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", controllers.CreateUser)
			auth.POST("/login", controllers.Login)
			auth.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
		}

		todos := v1.Group("/todos")
		todos.Use(middlewares.CheckAuth)
		{
			todos.POST("", controllers.CreateTodo)
			todos.GET("", controllers.GetAllTodos)
			todos.GET("/:id", controllers.GetTodo)
			todos.PUT("/:id", controllers.UpdateTodo)
			todos.DELETE("/:id", controllers.DeleteTodo)
		}

		tags := v1.Group("/tags")
		tags.Use(middlewares.CheckAuth)
		{
			tags.POST("", controllers.CreateTag)
			tags.GET("", controllers.GetAllTags)
			tags.GET("/:id", controllers.GetTag)
			tags.PUT("/:id", controllers.UpdateTag)
			tags.DELETE("/:id", controllers.DeleteTag)
		}
	}
	v1.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run route whenever triggered
	if err := route.Run(); err != nil {
		fmt.Printf("failed to start the server: %s", err)
		return
	}
}
