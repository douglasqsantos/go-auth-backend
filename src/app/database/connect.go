package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(){
	dsn := "db_user:db_password@tcp(mysql:3306)/db_go_api_template?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}
}