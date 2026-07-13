package v1

import (
	"api-for-shops-on-instagram/internal/module/instagram/delivery/http"

	"github.com/gin-gonic/gin"
)

func MetaInstagramRouter(groupV1 *gin.RouterGroup, handler *http.MetaInstagramHandler) {
	metaInstagram := groupV1.Group("/instagram")
	{
		metaInstagram.GET("/", handler.MetaInstagramInfo)
		metaInstagram.POST("/media", handler.Media)
		metaInstagram.GET("/conversations", handler.Conversation)
		metaInstagram.GET("/:conversation_id/messages", handler.MessageList)
		metaInstagram.GET("/message/:message_id/", handler.Message)
		metaInstagram.POST("/messages", handler.SendMessage)
	}
}
