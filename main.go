package main

import (
	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/routes"
	"gorm.io/gorm"
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
	var DB *gorm.DB
	config.ConnectDB()
	defer config.DisconnectDB(DB)

	// run all routes
	routes.Routes()
}
