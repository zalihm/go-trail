package main

import (
	"mlabs/goserve/controller"
	"mlabs/goserve/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/list", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}
