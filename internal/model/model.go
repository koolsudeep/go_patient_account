package model

import "github.com/jinzhu/gorm"

type PatientAccount struct {
	gorm.Model
	Name    string
	Email   string
	Age     int
	Gender  string
	Phone   string
	Address string
}

// Set the table name to "patient_accounts"
func (PatientAccount) TableName() string {
	return "patient_account"
}
