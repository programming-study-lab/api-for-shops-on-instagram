package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type InstagramConfig struct {
	GraphVersion string `json:"graphVersion"`
	Api          string `json:"api"`
	InstagramId  string `json:"instagramId"`
	AccessToken  string `json:"accessToken"`
}

func LoadInstagramConfig() *InstagramConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("เกิดข้อผิดพลาดในการโหลดข้อมูล Instagram : ", err.Error())
	}

	return &InstagramConfig{
		GraphVersion: os.Getenv("INSTAGRAM_GRAPH_VERSION"),
		Api:          os.Getenv("INSTAGRAM_API"),
		InstagramId:  os.Getenv("INSTAGRAM_ID"),
		AccessToken:  os.Getenv("INSTAGRAM_ACCESS_TOKEN"),
	}

}
