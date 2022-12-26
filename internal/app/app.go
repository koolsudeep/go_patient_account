package app

import (
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/go_patient_account/internal/handler"
	"github.com/koolsudeep/go_patient_account/internal/router"
)

func SetupRoutes(r *gin.Engine) {
	router.SetupHelloWorldRoutes(r, &handler.HelloWorldHandler{})
	router.SetupPatientAccountRoutes(r, &handler.PatientAccountHandler{})
}
