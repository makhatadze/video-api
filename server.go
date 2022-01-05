package main

import (
	"github.com/gin-gonic/gin"
	"vito-video/controller"
	"vito-video/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
