package config

import (
	"fmt"
	"os"

	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects go to the database
func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Kyiv",
		dbHost, dbUser, dbPass, dbName)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect Postgres database")
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		fmt.Printf("Failed to migrate: %v", err)
	}

	return db
}

// DisconnectDB is stopping connection to Postgres database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
