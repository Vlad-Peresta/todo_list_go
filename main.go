package main

import (
	"fmt"
	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/routes"
	"github.com/Vlad-Peresta/todo_list_go/src/models"
)

//	@title			Todo List API
//	@version		1.0.0
//	@description	API for working with the Todo List.
//	@contact.name	Vladyslav Peresta
//	@contact.url	https://github.com/Vlad-Peresta
//	@contact.email	perestavlad@gmail.com
//	@host			localhost:8080
//	@BasePath		/api/v1
func main() {
	config.ConnectDB()

	if err := config.DB.AutoMigrate(&models.Todo{}); err != nil {
		fmt.Printf("Failed to migrate: %v", err)
	}

	defer config.DisconnectDB(config.DB)

	// run all routes
	routes.Routes()
}
