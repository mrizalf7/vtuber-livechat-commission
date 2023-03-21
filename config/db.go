package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/models"

)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://rizal:root@localhost:5432/postgres"), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&models.PriceList{})
	db.AutoMigrate(&models.Commission{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.User{})
	DB = db
}