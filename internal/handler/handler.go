package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koolsudeep/go_patient_account/internal/database"
	"github.com/koolsudeep/go_patient_account/internal/model"
	"net/http"
)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(c *gin.Context) {
	c.String(200, "Hello, world!")
}

type PatientAccountHandler struct{}

func (h *PatientAccountHandler) GetPatientAccounts(c *gin.Context) {
	var patientAccounts []model.PatientAccount
	if err := database.DB.Find(&patientAccounts).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(http.StatusOK, patientAccounts)
}

func (h *PatientAccountHandler) GetPatientAccount(c *gin.Context) {
	var patientAccount model.PatientAccount
	patientAccountID := c.Param("id")
	if err := database.DB.Where("id = ?", patientAccountID).First(&patientAccount).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(http.StatusOK, patientAccount)
}

func (h *PatientAccountHandler) CreatePatientAccount(c *gin.Context) {
	var patientAccount model.PatientAccount
	if err := c.ShouldBindJSON(&patientAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&patientAccount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patientAccount)
}

func (h *PatientAccountHandler) UpdatePatientAccount(c *gin.Context) {
	var patientAccount model.PatientAccount
	patientAccountID := c.Param("id")
	if err := database.DB.Where("id = ?", patientAccountID).First(&patientAccount).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	err := c.BindJSON(&patientAccount)
	if err != nil {
		return
	}
	if err := database.DB.Save(&patientAccount).Error; err != nil {
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, patientAccount)
}

func (h *PatientAccountHandler) DeletePatientAccount(c *gin.Context) {
	var patientAccount model.PatientAccount
	patientAccountID := c.Param("id")
	if err := database.DB.Where("id = ?", patientAccountID).First(&patientAccount).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	if err := database.DB.Delete(&patientAccount).Error; err != nil {
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, gin.H{"id #" + patientAccountID: "deleted"})
}
