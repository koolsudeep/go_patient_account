package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/patientAccount/models"

	//"github.com/koolsudeep/patientAccount/models"
	"github.com/koolsudeep/patientAccount/router"
)

func main() {
	fmt.Println("Starting Patient API")
	r := gin.Default()

	router.InitRoutes(r)

	defer models.ConnectDatabase()
}
