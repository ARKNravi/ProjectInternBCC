package initializer

import (
	"ProjectBuahIn/buah"
	"ProjectBuahIn/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB 2345")
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&buah.Order{})
	DB.AutoMigrate(&buah.Buah{})
	DB.AutoMigrate(&buah.Cart{})
	DB.AutoMigrate(&buah.Order{})

}
