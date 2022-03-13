package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/patientAccount/models"
	"net/http"
)

func GetAllPatient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all Patients",
		"data":    []models.Patient{},
	})

}

func CreatePatient(c *gin.Context) {
	patient := models.Patient{}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.
		c.JSON(http.StatusOK, gin.H{
		"message": "Created Successfully",
		"data":    patient,
	})
}
func UpdatePatient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all Patients",
	})

}
func DeletePatient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all Patients",
	})

}
