package http

import (
	infrastructure "api-for-shops-on-instagram/internal/infrastructure/server/http/data_transfer_object"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// สำหรับการ Request ข้อมูลจาก API ภายนอก
type MetaInstagramRequest struct {
	instagramApi          string
	instagramGraphVersion string
	instagramId           string
	instagramAccessToken  string
}

// สำหรับการ Request ข้อมูลจาก API ภายนอก
func NewMetaInstagramRequest(instagramApi string, instagramGraphVersion string, instagramId string, instagramAccessToken string) *MetaInstagramRequest {
	return &MetaInstagramRequest{
		instagramApi:          instagramApi,
		instagramGraphVersion: instagramGraphVersion,
		instagramId:           instagramId,
		instagramAccessToken:  instagramAccessToken,
	}
}

// ดึงข้อมูล เช่น Instagram ID ของตนเอง
func (metaInstagramRequest *MetaInstagramRequest) InstagramGetInfo() (*map[string]any, error) {

	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/me"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	instagramResponse, err := http.Get(url)

	if err != nil {
		log.Fatal("meta_instagram_request.go(InstagramGetInfo) ดึงข้อมูลจาก API Meta Instagram ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(instagramResponse.Body)
	if err != nil {
		log.Fatal("meta_instagram_request.go(InstagramGetInfo) แปลงข้อมูลเป็น JSON: ", err.Error())
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("meta_instagram_request.go(InstagramGetInfo): เปลี่ยน JSON เป็น map: ", err.Error())
		return nil, err
	}

	return &result, nil

}

// สร้าง container สำหรับการเตรียมข้อมูล ก่อนจะโพสต์ Instagram
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

// โพสต์ Content ออกสู่สาธารณะ จาก container ที่ได้สร้างไว้
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMediaPublish(creationId *map[string]any) (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + metaInstagramRequest.instagramId
	url += "/media_publish"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	id := map[string]string{}

	createionId, ok := (*creationId)["id"].(string)
	if ok == true {
		id["creation_id"] = createionId
	}

	metaRequest, err := json.Marshal(id)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMediaPublish) แปลงข้อมูล จาก map[string]any เป็น Marshal: ", err.Error())
		return nil, err
	}

	// หน่วงเวลาไว้ 1 วินาที เพื่อไม้ให้ api ของ Meta Instagram ปฏิเสธการร้องขอ
	// เพราะ หากร้องขอเร็วเกินไปจะถูก Meta Instagram ปฏิเสธ
	delay := 1
	time.Sleep(time.Duration(delay) * time.Second)

	bufferMetaRequest := bytes.NewBuffer(metaRequest)
	metaInstagramResponse, err := http.Post(url, "application/json", bufferMetaRequest)
	if err != nil {
		log.Fatalln("metaInstagramRequest.go(MetaInstagramMediaPublish) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMediaPublish) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramMediaPublish) แปลงข้อมูลจาก json เป็น result: ", err.Error())
		return nil, err
	}

	return &result, nil

}

// ดึงข้อมูล การสนทนา เพื่อรับ conversation_id
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramConversations() (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + metaInstagramRequest.instagramId
	url += "/conversations"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	metaInstagramResponse, err := http.Get(url)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramConversations) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramConversations) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramConversations) แปลงข้อมูลจาก json เป็น result: ", err.Error())
		return nil, err
	}

	return &result, nil

}

// ดึงจ้อมูลการสนทนา โดยใช้ conversation_id เพื่อรับ message_id
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMessageList(ctx *gin.Context) (*map[string]any, error) {

	conversation_id := ctx.Param("conversation_id")

	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + conversation_id
	url += "/messages"
	url += "?fields=" + "id,form,updated_time,messages"
	url += "&access_token=" + metaInstagramRequest.instagramAccessToken

	metaInstagramResponse, err := http.Get(url)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMessageList) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMessageList) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramMessageList) แปลงข้อมูลจาก json เป็น result: ", err.Error())
		return nil, err
	}

	return &result, nil
}

// ดึงข้อความแชท โดยใช้ message_id หมายเหตุ ข้อมูลที่ได้มี instagram_id ของ ผู้ใช้งาน Instagram ของคู่สนทนาด้วย
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMessage(ctx *gin.Context) (*map[string]any, error) {
	message_id := ctx.Param("message_id")

	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + message_id
	url += "?fields=" + "id,created_time,from,to,message"
	url += "&access_token=" + metaInstagramRequest.instagramAccessToken

	metaInstagramResponse, err := http.Get(url)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaINstagramMessage) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramMessage) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagrmMesssage) แปลงข้อมูลจาก json เป็น result: ", err.Error())
		return nil, err
	}

	return &result, nil
}

// ส่งข้อความ โดยใช้ instagram_id ของคู่สนทนา หมายเหตุ instagram_id มาพร้อมกับข้อความแชท
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramSendMessage(sendMessageModel *infrastructure.MetaInstagramSendMessageModel) (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + metaInstagramRequest.instagramId
	url += "/messages"
	url += "?access_token=" + metaInstagramRequest.instagramAccessToken

	metaRequest, err := json.Marshal(sendMessageModel)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramSendMessage) แปลงข้อมูล จาก map[string]any เป็น Marshal: ", err.Error())
		return nil, err
	}

	bufferMetaRequest := bytes.NewBuffer(metaRequest)
	metaInstagramResponse, err := http.Post(url, "application/json", bufferMetaRequest)
	if err != nil {
		log.Fatalln("metaInstagramRequest.go(MetaInstagramSendMessage) ส่งคำขอไปยัง Meta Instagram: ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(metaInstagramResponse.Body)
	if err != nil {
		log.Fatalln("meta_instagram_request.go(MetaInstagramSendMessage) แปลงข้อมูล ที่ได้รับมาเป็น Byte: ", err.Error())
		return nil, err
	}

	result := make(map[string]any)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("meta_instagram_request(MetaInstagramSendMessage) แปลงข้อมูลจาก json เป็น result: ", err.Error())
		return nil, err
	}

	return &result, nil
}

// ดึงรายการโพส์ของ Instagram ที่เคยโพสต์ หมายเหตุ ข้อมูลที่ได้รับมี media_id สามารถนำไปดูรายละเอียดโพสต์
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMediaList() (*map[string]any, error) {
	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/me"
	url += "/media"
	url += "?fields=" + "id,caption,media_type,media_url,permalink,timestamp"
	url += "&access_token=" + metaInstagramRequest.instagramAccessToken

	instagramResponse, err := http.Get(url)

	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaList) ดึงข้อมูลจาก API Meta Instagram ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(instagramResponse.Body)
	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaList) แปลงข้อมูลเป็น JSON: ", err.Error())
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaList): เปลี่ยน JSON เป็น map: ", err.Error())
		return nil, err
	}

	return &result, nil
}

// ใช้ media_id ในการดึงข้อมูลรายละเอียดโพสต์
func (metaInstagramRequest *MetaInstagramRequest) MetaInstagramMediaDetail(ctx *gin.Context) (*map[string]any, error) {

	mediaId := ctx.Param("media_id")

	url := metaInstagramRequest.instagramApi
	url += "/" + metaInstagramRequest.instagramGraphVersion
	url += "/" + mediaId
	url += "?fields=" + "id,caption,media_type,media_url,permalink,timestamp"
	url += "&access_token=" + metaInstagramRequest.instagramAccessToken

	instagramResponse, err := http.Get(url)

	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaDetail) ดึงข้อมูลจาก API Meta Instagram ", err.Error())
		return nil, err
	}

	body, err := io.ReadAll(instagramResponse.Body)
	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaDetail) แปลงข้อมูลเป็น JSON: ", err.Error())
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("meta_instagram_request.go(MetaInstagramMediaDetail): เปลี่ยน JSON เป็น map: ", err.Error())
		return nil, err
	}

	return &result, nil
}
