package http

import (
	server "api-for-shops-on-instagram/internal/infrastructure/server/http"
	infrastructure "api-for-shops-on-instagram/internal/infrastructure/server/http/data_transfer_object"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetaInstagramHandler struct {
	instagramApi          string
	instagramGraphVersion string
	instagramId           string
	instagramAccessToken  string
	metaInstagramRequest  *server.MetaInstagramRequest
	// usecase *usecase.
}

// func NewMetaInstagramHandler(instagramApi string, instagramGraphVersion string, instagramId string, instagramAccessToken string) *MetaInstagramHandler {
func NewMetaInstagramHandler(metaInstagramMediaRequest *server.MetaInstagramRequest) *MetaInstagramHandler {
	// metaInstagramUsecase := usecase.NewMetaInstagramUsecase(&newMetaInstagramHandler)
	return &MetaInstagramHandler{
		metaInstagramRequest: metaInstagramMediaRequest,
	}
}

func (metaIG *MetaInstagramHandler) MetaInstagramInfo(ctx *gin.Context) {
	metaInstagramInfo, err := metaIG.metaInstagramRequest.InstagramGetInfo()
	if err != nil {
		log.Fatalln("meta_instagram_http.go(Conversation) รับข้อมูลจาก Meta Instagram: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &metaInstagramInfo,
		},
	)
}

func (metaIG *MetaInstagramHandler) Media(ctx *gin.Context) {
	createionId, err := metaIG.metaInstagramRequest.MetaInstagramMedia(ctx)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(Media) รับค่าจาก MetaInstagramMedia: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	response, err := metaIG.metaInstagramRequest.MetaInstagramMediaPublish(createionId)

	if err != nil {
		log.Fatalln("meta_instagram_http.go(Media) รับค่าจาก MetaInstagramMediaPublish: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  true,
				Message: "success",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "success",
			Data:    response,
		},
	)

}

func (metaIG *MetaInstagramHandler) Conversation(ctx *gin.Context) {
	conversation, err := metaIG.metaInstagramRequest.MetaInstagramConversations()
	if err != nil {
		log.Fatalln("meta_instagram_http.go(Conversation) รับข้อมูลจาก Meta Instagram: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &conversation,
		},
	)

}

func (metaIG *MetaInstagramHandler) MessageList(ctx *gin.Context) {
	messageList, err := metaIG.metaInstagramRequest.MetaInstagramMessageList(ctx)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(MessageList) รับข้อมูลจาก MetaInstagramMessageList: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &messageList,
		},
	)

}

func (metaIG *MetaInstagramHandler) Message(ctx *gin.Context) {

	message, err := metaIG.metaInstagramRequest.MetaInstagramMessage(ctx)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(Message) รับข้อมูลจาก MetaInstagramMessage: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &message,
		},
	)

}

func (metaIG *MetaInstagramHandler) SendMessage(ctx *gin.Context) {
	sendMessageModel := infrastructure.MetaInstagramSendMessageModel{}

	err := ctx.ShouldBindBodyWithJSON(&sendMessageModel)
	if err != nil {
		log.Fatalln("meta_instagram_http(SendMessage) รับข้อมูลจาก JSON: ", err.Error())
		return
	}

	response, err := metaIG.metaInstagramRequest.MetaInstagramSendMessage(&sendMessageModel)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(SendMessage) รับข้อมูลจาก MetaInstagramSendMessage: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &response,
		},
	)
}

func (metaIG *MetaInstagramHandler) MediaList(ctx *gin.Context) {
	mediaList, err := metaIG.metaInstagramRequest.MetaInstagramMediaList()
	if err != nil {
		log.Fatalln("meta_instagram_http.go(MediaList) รับข้อมูลจาก MetaInstagramMediaList: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &mediaList,
		},
	)
}

func (metaIG *MetaInstagramHandler) MediaDetail(ctx *gin.Context) {
	mediaDetail, err := metaIG.metaInstagramRequest.MetaInstagramMediaDetail(ctx)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(MediaDetail) รับข้อมูลจาก MetaInstagramMediaDetail: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			infrastructure.APIResponse{
				Status:  false,
				Message: "error",
				Data:    err.Error(),
			},
		)
		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		infrastructure.APIResponse{
			Status:  true,
			Message: "succsss",
			Data:    &mediaDetail,
		},
	)
}
