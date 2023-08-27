package Database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {

	dsn := "root:@tcp(127.0.0.1:3306)/go-upload-cloudinary-gin?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	DB = db

	fmt.Println("Database Connection Success")
}