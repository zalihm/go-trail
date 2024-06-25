package controller

import (
	"mlabs/goserve/entity"
	"mlabs/goserve/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return controller.service.Save(video)
}
