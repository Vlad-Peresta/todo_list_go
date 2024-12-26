package main

import (
	config "github.com/Vlad-Peresta/todo_list_go/src/conf"
	"github.com/Vlad-Peresta/todo_list_go/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	// run all routes
	routes.Routes()
}
