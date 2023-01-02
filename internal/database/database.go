package database

import (
	"encoding/json"
	"fmt"
	"github.com/koolsudeep/go_patient_account/internal/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

// Init connects to the database and initializes the DB global variable
func Init() {
	// Load the environment variables from the .env file
	errLoad := godotenv.Load(".env")
	if errLoad != nil {
		fmt.Println("having problem here", errLoad)
		log.Fatal("Error loading .env file")
	}

	// Read the database connection details from the environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	// Connect to the database
	var errOpen error
	DB, errOpen = gorm.Open("postgres", connectionString)
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	//// Already created the table, so don't need following:
	//DB.AutoMigrate(&model.PatientAccount{})

	// Ping the database to check the connection
	if errPing := DB.DB().Ping(); errPing != nil {
		log.Fatal(errPing)
	}
	fmt.Println("Successfully connected to database!")

	patientAccount := model.PatientAccount{
		Name:    "ram hari",
		Email:   "ram@example.com",
		Age:     40,
		Gender:  "male",
		Phone:   "623-456-7890",
		Address: "8585 Main Street,Irving,USA",
	}

	patientAccountJSON, err := json.Marshal(patientAccount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(patientAccountJSON))
}
