package v1

import (
	"api-for-shops-on-instagram/internal/module/instagram/delivery/http"

	"github.com/gin-gonic/gin"
)

func MetaInstagramRouter(groupV1 *gin.RouterGroup, handler *http.MetaInstagramHandler) {
	metaInstagram := groupV1.Group("/instagram")
	{
		metaInstagram.GET("/", handler.MetaInstagramInfo)                    // ตรวจสอบข้อมูลบัญชี Instagram ของตัวเอง
		metaInstagram.POST("/media", handler.Media)                          // โพสต์ รูปภาพ และ caption โพสต์รูปโดย URL ของรูปภาพ
		metaInstagram.GET("/conversations", handler.Conversation)            // รายการ การสนทนา หมายเหตุ สามารถดู conversation_id ได้
		metaInstagram.GET("/:conversation_id/messages", handler.MessageList) // แสดงรายการ การสนทนา หมายเหตุ สามารถดู message_id ได้
		metaInstagram.GET("/message/:message_id/", handler.Message)          // แสดงข้อความในแชท
		metaInstagram.POST("/messages", handler.SendMessage)                 // ส่งข้อความ
		metaInstagram.GET("/media", handler.MediaList)                       // แสดงรายการ content ที่อยู่ใน โปรไฟล์
		metaInstagram.GET("/media/:media_id", handler.MediaDetail)           // แสดงรายละเอียด content ที่อยู่ในโปรไฟล์
	}
}
