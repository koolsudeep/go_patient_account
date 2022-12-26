package main

import (
	"fmt"
	"github.com/koolsudeep/go_patient_account/internal/database"
	"github.com/koolsudeep/go_patient_account/internal/model"
	"log"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/go_patient_account/internal/app"
)

func main() {
	patientAccount := model.PatientAccount{
		Name:    "John Smith",
		Email:   "john@example.com",
		Age:     30,
		Gender:  "male",
		Phone:   "123-456-7890",
		Address: "123 Main Street,Irving,USA",
	}

	patientAccountJSON, err := json.Marshal(patientAccount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(patientAccountJSON))
	database.Init()
	r := gin.Default()
	app.SetupRoutes(r)
	log.Fatal(r.Run())

}
