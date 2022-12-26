package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	var errOpen error
	DB, errOpen = gorm.Open("postgres", connectionString)
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	if errPing := DB.DB().Ping(); errPing != nil {
		log.Fatal(errPing)
	}
	
	fmt.Println("Successfully connected to database!")
}
