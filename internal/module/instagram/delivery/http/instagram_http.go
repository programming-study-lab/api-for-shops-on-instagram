package http

// import (
// 	infrastructure "api-for-shops-on-instagram/internal/infrastructure/server/http/data_transfer_object"
// 	"api-for-shops-on-instagram/internal/module/instagram/usecase"

// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type InstagramHandler struct {
// 	usecase usecase.InstagramUsecase
// }

// func NewInstagramHandler(usecase usecase.InstagramUsecase) *InstagramHandler {
// 	return &InstagramHandler{
// 		usecase: usecase,
// 	}
// }

// func (instagram *InstagramHandler) GetInfo(ctx *gin.Context) {

// 	instagramInfo, err := instagram.usecase.InstagramGetInfo(ctx)

// 	if err != nil {
// 		log.Fatal("instagram_http(GetInfo): ", err.Error())
// 		ctx.AbortWithStatusJSON(
// 			http.StatusBadRequest,
// 			gin.H{
// 				"error": "Error",
// 			},
// 			// ig.response.ResponseError("error", nil),
// 		)
// 		return
// 	}

// 	resDTO := infrastructure.APIResponse{
// 		Status:  true,
// 		Message: "success",
// 		Data:    &instagramInfo,
// 	}

// 	ctx.AbortWithStatusJSON(
// 		http.StatusOK,
// 		resDTO,
// 	)

// }
