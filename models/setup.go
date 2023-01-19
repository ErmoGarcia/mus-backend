package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {

		DbHost := os.Getenv("DB_HOST")
		DbUser := os.Getenv("DB_USER")
		DbPassword := os.Getenv("DB_PASSWORD")
		DbName := os.Getenv("DB_NAME")
		DbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf("host=%s user%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Madrid", DbHost, DbUser, DbPassword, DbName, DbPort)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("Cannot connect to Postgress database.")
			log.Fatal("connection error:", err)
		} else {
			fmt.Println("We are connected to the Postgress production database.")
		}

		db.AutoMigrate(&User{})

	} else {

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			fmt.Println("Cannot connect to SQLite database.")
			log.Fatal("connection error:", err)
		} else {
			fmt.Println("We are connected to the SQLite testing database.")
		}

		db.AutoMigrate(&User{})

	}
}
