package main

import (
	"github.com/koolsudeep/go_patient_account/internal/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/go_patient_account/internal/app"
)

func main() {

	//handler := &app.Handler{DB: database.DB}
	//router.SetupPatientAccountRoutes(r, handler)

	r := gin.Default()
	app.SetupRoutes(r)
	log.Fatal(r.Run())
	database.Init()
}
