package config

import (
	"test/go/models"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	Connect()
}

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
	color.Green("Connection Opened to Database... Success!")

	db.AutoMigrate(&models.User{})
	color.Green("Database Migrated... Success!")
}

func GetDB() *gorm.DB {
	return db
}
