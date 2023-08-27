package Migrations

import (
	"fmt"
	"Go-Upload-Cloudinary-Gin/Models"
	"Go-Upload-Cloudinary-Gin/Database"
)

func Migration() {
	err := Database.DB.AutoMigrate(
		&Models.User{},
	)

	if err != nil {
		fmt.Println("Can't Running Migrations")
		return
	}

	fmt.Println("Migrations Success")
}