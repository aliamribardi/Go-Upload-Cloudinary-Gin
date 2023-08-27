package main

import (
	"Go-Upload-Cloudinary-Gin/Database/Migrations"
	"Go-Upload-Cloudinary-Gin/Database"
	"fmt"
)

func main() {
	Database.ConnectionDatabase()
	Migrations.Migration()
	fmt.Println("Success run the program")
}