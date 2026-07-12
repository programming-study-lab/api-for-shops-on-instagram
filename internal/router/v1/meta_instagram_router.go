package v1

import (
	"api-for-shops-on-instagram/internal/module/instagram/delivery/http"

	"github.com/gin-gonic/gin"
)

func MetaInstagramRouter(groupV1 *gin.RouterGroup, handler *http.MetaInstagramHandler) {
	metaInstagram := groupV1.Group("/instagram")
	{
		metaInstagram.POST("/media", handler.Media)
	}
}
