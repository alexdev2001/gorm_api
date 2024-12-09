package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm_api/models"
	"log"
	"os"
)

func DatabaseConnection() {
	// load the .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get the values from the .env
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	// format the values to a connection string
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s", db_host, db_name, db_port, db_user, db_password)

	// open the database connection

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database")
	}

	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatal("failed to migrate the database")
	}

}
