package main

import (
	"github.com/Vlad-Peresta/todo_list_go/src/config"
	"github.com/Vlad-Peresta/todo_list_go/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConectDB()
)

func main() {
	defer config.DisconnectDB(db)

	// run all routes
	routes.Routes()
}
