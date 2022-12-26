package router

import (
	"github.com/gin-gonic/gin"
)

type HelloWorldHandler interface {
	HelloWorld(c *gin.Context)
}

func SetupHelloWorldRoutes(r *gin.Engine, handler HelloWorldHandler) {
	r.GET("/hello", handler.HelloWorld)
}

type PatientAccountHandler interface {
	GetPatientAccounts(c *gin.Context)
	GetPatientAccount(c *gin.Context)
	CreatePatientAccount(c *gin.Context)
	UpdatePatientAccount(c *gin.Context)
	DeletePatientAccount(c *gin.Context)
}

func SetupPatientAccountRoutes(r *gin.Engine, handler PatientAccountHandler) {
	r.GET("/patientAccounts", handler.GetPatientAccounts)
	r.GET("/patientAccounts/:id", handler.GetPatientAccount)
	r.POST("/patientAccounts", handler.CreatePatientAccount)
	r.PUT("/patientAccounts/:id", handler.UpdatePatientAccount)
	r.DELETE("/patientAccounts/:id", handler.DeletePatientAccount)
}
