package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model  `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	DateOfBirth string    `json:"date_of_birth"`
	Phone       []Phone   `json:"phone"`
	Email       []Email   `json:"email"`
	Address     *Address  `json:"address"`
	Disease     []Disease `json:"disease"`
	UserId      uint      `json:"user_id"`
}

type Phone struct {
	gorm.Model
	PrimaryPhone   uint64 `json:"primary_phone"`
	SecondaryPhone uint64 `json:"secondary_phone"`
	PatientID      uint
}

type Email struct {
	gorm.Model
	PrimaryEmail   string `json:"primary_email"`
	SecondaryEmail string `json:"secondary_email"`
	PatientID      uint
}

type Address struct {
	gorm.Model
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	PatientID uint
}

type Disease struct {
	gorm.Model
	Illness   string `json:"illness"`
	Symptom   string `json:"symptom"`
	PatientID uint
}
