package database

import (
	"github.com/harish1907/go-udemy/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	dsn := "host=mel.db.elephantsql.com user=xkgvvbow password=5HMAsTxfAbp1k65djA3N8Y8PGruT8BTA dbname=xkgvvbow port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection error..")
	}
	DB.AutoMigrate(&models.MyUser{}, &models.Roles{})
}
