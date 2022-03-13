package router

import (
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/patientAccount/controllers"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/", controllers.Welcome)
	r.GET("/ping", controllers.Ping)

	//patients endpoints
	r.GET("/patients", controllers.GetAllPatient)
	r.POST("/patient", controllers.CreatePatient)
	r.PUT("/patient", controllers.UpdatePatient)
	r.DELETE("/patient", controllers.DeletePatient)

	r.Run()

}
