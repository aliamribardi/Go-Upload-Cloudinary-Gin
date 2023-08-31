package Controllers

import (
	"log"
	"context"
	"os"
	"Go-Upload-Cloudinary-Gin/Database"
	"Go-Upload-Cloudinary-Gin/Models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UserIndex(c *gin.Context) {

}

func UserUpload(c *gin.Context) {
	var user Models.User
	
	godotenv.Load()

	urlCloudinary := os.Getenv("CLOUDINARY_URL")

	fileImage, err := c.FormFile("Image")

	image, _ := fileImage.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(urlCloudinary)

	resp, _ := cldService.Upload.Upload(ctx, image, uploader.UploadParams{})

	log.Println(resp.SecureURL)

	user.Image = resp.SecureURL

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to Upload"})
		return
	}
	
	Database.DB.Create(&user)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": user,
	})
	
}