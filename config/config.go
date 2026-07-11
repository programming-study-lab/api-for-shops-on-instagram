package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppURL     string `json:"appURL"`
	AppPort    string `json:"appPort"`
	AppAddress string `json:"appAddress"`
}

func LoadAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("config(LoadAppConfig): โหลดข้อมูลล้มเหลว:  ", err.Error())
	}

	return &AppConfig{
		AppURL:     os.Getenv("APP_URL"),
		AppPort:    os.Getenv("APP_PORT"),
		AppAddress: os.Getenv("APP_ADDRESS"),
	}
}
