package db

import (
	"fmt"
	"log"
	"os"

	"github.com/ErmoGarcia/mus-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupPostgres() (*gorm.DB, error) {

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Madrid",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Cannot connect to Postgress database.")
		log.Fatal("connection error:", err)
	}

	fmt.Println("We are connected to the Postgress production database.")
	return db, nil
}

func setupSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to SQLite database.")
		log.Fatal("connection error:", err)
	}

	fmt.Println("We are connected to the SQLite development database.")
	return db, err
}

func ConnectDataBase() error {

	var db *gorm.DB
	var err error

	mode := os.Getenv("GIN_MODE")

	switch mode {
	case "release":
		db, err = setupPostgres()
	default:
		db, err = setupSQLite()
	}

	if err != nil {
		return err
	}

	err = models.AutoMigrate(db)

	if err != nil {
		return err
	}

	dbInstance = db

	return nil
}
