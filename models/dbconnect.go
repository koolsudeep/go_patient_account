package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&Patient{},
		&Phone{},
		&Email{},
		&Disease{},
	)
	if err != nil {
		panic(err)
	}
	return db

}
