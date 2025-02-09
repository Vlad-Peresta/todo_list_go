package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define database client
var DB *gorm.DB

// ConnectDB connects go to the database
func ConnectDB() {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("failed to load env file")
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
		panic("failed to connect Postgres database")
	}

	DB = db
}

// DisconnectDB is stopping connection to Postgres database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to kill connection from database")
	}
	dbSQL.Close()
}
