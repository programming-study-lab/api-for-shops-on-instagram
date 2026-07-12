package http

import (
	infrastructure "api-for-shops-on-instagram/internal/infrastructure/server/http/data_transfer_object"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MetaInstagramRequest struct {
	// ctx                   *gin.Context
	instagramApi          string
	instagramGraphVersion string
	instagramId           string
	instagramAccessToken  string
}

func NewMetaInstagramRequest(instagramApi string, instagramGraphVersion string, instagramId string, instagramAccessToken string) *MetaInstagramRequest {
	return &MetaInstagramRequest{
		instagramApi:          instagramApi,
		instagramGraphVersion: instagramGraphVersion,
		instagramId:           instagramId,
		instagramAccessToken:  instagramAccessToken,
	}
}

func (metaInstagramRequest *MetaInstagramRequest) FetchInstagramInfo() {

}

func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMedia(ctx *gin.Context) (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + metaInstagramRequest.instagramId
	url += "/media"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	metaInstagramMediaRequest := &infrastructure.MetaInstagramMediaRequest{}

	err := ctx.ShouldBindBodyWithJSON(&metaInstagramMediaRequest)

	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMedia) รับข้อมูลจาก Request: ", err.Error())
		return nil, err
	}

	metaRequest, err := json.Marshal(&metaInstagramMediaRequest)
	fmt.Printf("\ntest: %s\n", string(metaRequest))
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMedia) แปลงข้อมูล จาก Request เป็น Marshaler: ", err.Error())
		return nil, err
	}

	bufferMetaRequest := bytes.NewBuffer(metaRequest)
	metaInstagramResponse, err := http.Post(url, "application/json", bufferMetaRequest)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMedia) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	fmt.Printf("\nbody: %s\n", string(body))
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMedia) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramMedia) แปลงข้อมูลจาก json เป็น MetaInstagramMediaResponse: ", err.Error())
		return nil, err
	}

	return &result, nil

}

func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMediaPublish(creationId *map[string]any) (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + metaInstagramRequest.instagramId
	url += "/media_publish"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	// fmt.Printf("\n\n[url] %s\n\n", url)

	id := map[string]string{}

	createionId, ok := (*creationId)["id"].(string)
	if ok == true {
		// id := map[string]string{
		// "creation_id": createionId,
		// }
		id["creation_id"] = createionId
	}

	// id := map[string]string{
	// 	"creation_id": creationId.Id,
	// }

	metaRequest, err := json.Marshal(id)
	fmt.Printf("\ntest: %s\n", string(metaRequest))
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMediaPublish) แปลงข้อมูล จาก map[string]any เป็น Marshal: ", err.Error())
		return nil, err
	}

	delay := 1
	time.Sleep(time.Duration(delay) * time.Second)

	bufferMetaRequest := bytes.NewBuffer(metaRequest)
	metaInstagramResponse, err := http.Post(url, "application/json", bufferMetaRequest)
	if err != nil {
		log.Fatalln("metaInstagramRequest.go(MetaInstagramMediaPublish) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	fmt.Printf("\nbody: %s\n", string(body))
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMediaPublish) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramMediaPublish) แปลงข้อมูลจาก json เป็น MetaInstagramMediaResponse: ", err.Error())
		return nil, err
	}

	// fmt.Println("[err] ", result)
	return &result, nil

}
