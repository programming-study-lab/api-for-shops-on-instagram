package usecase

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// )

// type instagramInfoUsecase struct {
// 	instagramApi          string
// 	instagramGraphVersion string
// 	instagramAccessToken  string
// }

// func NewInstagramInfoUsecase(instagramApi string, instagramGraphVersion string, instagramAccessToken string) InstagramUsecase {
// 	return &instagramInfoUsecase{
// 		instagramApi:          instagramApi,
// 		instagramGraphVersion: instagramGraphVersion,
// 		instagramAccessToken:  instagramAccessToken,
// 	}
// }

// func (usecase *instagramInfoUsecase) InstagramGetInfo(ctx context.Context) (*[]map[string]any, error) {

// 	url := usecase.instagramApi
// 	url += "/" + usecase.instagramGraphVersion
// 	url += "me"
// 	url += "?access_token=" + usecase.instagramAccessToken

// 	instagramResponse, err := http.Get(url)

// 	if err != nil {
// 		log.Fatal("instagram_usecase(InstagramGetInfo): ", err.Error())
// 		return nil, err
// 	}

// 	body, err := io.ReadAll(instagramResponse.Body)
// 	if err != nil {
// 		log.Fatal("instagram_usecase(InstagramGetInfo): ", err.Error())
// 		return nil, err
// 	}

// 	var convertTypeInstagramResponse map[string]interface{}
// 	err = json.Unmarshal(body, &convertTypeInstagramResponse)
// 	if err != nil {
// 		log.Fatal("instagram_usecase.go(InstagramGetInfo): เปลี่ยน Json เป็น interface{}: ", err.Error())
// 		return nil, err
// 	}

// 	var instagramInfo []map[string]any

// 	instagramInfo = append(instagramInfo, convertTypeInstagramResponse)

// 	return &instagramInfo, nil
// }
