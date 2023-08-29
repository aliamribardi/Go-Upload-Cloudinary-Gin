package Routes

import (
	"Go-Upload-Cloudinary-Gin/Controllers"
	"github.com/gin-gonic/gin"
)

func Route() {
	Route := gin.Default()

	Route.POST("/api", Controllers.UserUpload)

	Route.Run("localhost:8080")
}