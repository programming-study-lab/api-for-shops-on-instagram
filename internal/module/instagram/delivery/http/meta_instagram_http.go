package http

import (
	server "api-for-shops-on-instagram/internal/infrastructure/server/http"
	"fmt"
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

	fmt.Println("[debug] %s", createionId)

	response, err := metaIG.metaInstagramRequest.MetaInstagramMediaPublish(createionId)
	// delay := 300
	// time.Sleep(time.Duration(delay) * time.Millisecond)
	if err != nil {
		log.Fatalln("meta_instagram_http.go(Media) รับค่าจาก MetaInstagramMediaPublish: ", err.Error())
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	// value, ok := (*response)["id"].(string)
	// if ok != true {
	// 	// log.Fatalln("meta_instagram_http.go(Media) response == nil: ", err.Error())
	// 	ctx.AbortWithStatusJSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"test": "ok",
	// 		},
	// 	)
	// 	return
	// }

	// fmt.Printf("\n[debug] %s, %s\n", response, value)
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "message",
			"data":    response,
		},
	)

}

// func (metaIG *MetaInstagramHandler) Media(ctx *gin.Context) {
// 	url := metaIG.instagramApi
// 	url += "/" + metaIG.instagramGraphVersion
// 	url += "/" + metaIG.instagramId
// 	url += "/media"
// 	url += "?access_token=" + metaIG.instagramAccessToken

// 	fmt.Printf("\n\n[url] %s\n\n", url)

// 	metaInstagramMediaRequest := &infrastructure.MetaInstagramMediaRequest{}

// 	err := ctx.ShouldBindBodyWithJSON(&metaInstagramMediaRequest)

// 	if err != nil {
// 		log.Fatalln("meta_instagram_http(Media) รับข้อมูลจาก Request: ", err.Error())
// 		// return nil, err
// 	}

// 	metaRequest, err := json.Marshal(&metaInstagramMediaRequest)
// 	fmt.Printf("\ntest: %s\n", string(metaRequest))
// 	if err != nil {
// 		log.Fatalln("meta_instagram_http(Media) แปลงข้อมูล จาก Request เป็น Marshaler: ", err.Error())
// 		// return nil, err
// 	}

// 	bufferMetaRequest := bytes.NewBuffer(metaRequest)
// 	metaInstagramResponse, err := http.Post(url, "application/json", bufferMetaRequest)
// 	if err != nil {
// 		log.Fatalln("meta_instagram_http(Media) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
// 		// return nil, err
// 	}

// 	body, err := io.ReadAll(metaInstagramResponse.Body)
// 	fmt.Printf("\nbody: %s\n", string(body))
// 	if err != nil {
// 		log.Fatalln("meta_instagram_http(media) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
// 		// return nil, err
// 	}

// 	// convertTypeInstagramResponse := infrastructure.MetaInstagramMediaResponse{}
// 	convertTypeInstagramResponse := make(map[string]any)
// 	err = json.Unmarshal(body, &convertTypeInstagramResponse)
// 	if err != nil {
// 		log.Fatalln("meta_instagram_http(media) แปลงข้อมูลจาก json เป็น MetaInstagramMediaResponse: ", err.Error())
// 		// return nil, err
// 	}

// 	// metaInstagramMediaResponse := []infrastructure.MetaInstagramMediaResponse{}
// 	var metaInstagramMediaResponse []map[string]any
// 	metaInstagramMediaResponse = append(metaInstagramMediaResponse, convertTypeInstagramResponse)

// 	ctx.AbortWithStatusJSON(
// 		http.StatusOK,
// 		metaInstagramMediaResponse,
// 	)

// 	// ctx.JSON(
// 	// 	http.StatusOK,
// 	// 	gin.H{
// 	// 		"test": metaInstagramResponse,
// 	// 	},
// 	// )

// 	// return nil, nil

// }
