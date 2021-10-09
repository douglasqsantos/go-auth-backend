package database

import (
	"app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "db_user:db_password@tcp(mysql:3306)/db_go_api_template?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
	if err != nil {
		return
	}
}
