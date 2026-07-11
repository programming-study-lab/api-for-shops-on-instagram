package v1

import (
	"github.com/gin-gonic/gin"
)

type ImageHttpHandler struct {
	// context *gin.Context
}

func NewImageHttp() *ImageHttpHandler {
	return &ImageHttpHandler{}
}

func (ImageHttpHandler) GetImage(ctx *gin.Context) {
	image_name := ctx.Param("image_name")
	ctx.File("assets/public/image/" + image_name)
}
