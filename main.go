package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/cloudinary-gin-api/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello from Cloudinary",
		})
	})

	router.POST("file", controllers.FileUpload())
	router.POST("/remote", controllers.RemoteUpload())

	router.Run()
}
