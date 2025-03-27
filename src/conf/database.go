package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define database client
var DB *gorm.DB

// ConnectDB connects go to the database
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("DB_PORT_INTERNAL")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

// DisconnectDB is stopping connection to Postgres database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to kill connection from database")
	}
	defer dbSQL.Close()
}
