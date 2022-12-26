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
