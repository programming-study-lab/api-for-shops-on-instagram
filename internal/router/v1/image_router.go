package v1

import (
	image "api-for-shops-on-instagram/internal/module/image/delivery/http"

	"github.com/gin-gonic/gin"
)

func ImageRouter(groupV1 *gin.RouterGroup, handler *image.ImageHttpHandler) {
	image := groupV1.Group("/image")
	{
		image.GET("/:image_name", handler.GetImage) // API รูปภาพ
	}
}
