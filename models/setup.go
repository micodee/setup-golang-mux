package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	db, err := gorm.Open(mysql.Open("admin:admin@tcp(localhost:3306)/gorilla_crud"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	DB = db

	fmt.Println("Connected to Database")
}
