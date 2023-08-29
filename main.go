package main

import (
	"Go-Upload-Cloudinary-Gin/Routes"
	"Go-Upload-Cloudinary-Gin/Database/Migrations"
	"Go-Upload-Cloudinary-Gin/Database"
	"fmt"
)

func main() {
	Database.ConnectionDatabase()
	Migrations.Migration()
	Routes.Route()
	fmt.Println("Success run the program")
}