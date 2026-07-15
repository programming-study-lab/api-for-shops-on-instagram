package main

import (
	"api-for-shops-on-instagram/config"
	server "api-for-shops-on-instagram/internal/infrastructure/server/http"
	image "api-for-shops-on-instagram/internal/module/image/delivery/http"
	http "api-for-shops-on-instagram/internal/module/instagram/delivery/http"
	"api-for-shops-on-instagram/internal/router"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// โหลดข้อมูลจาก .env
	igConfig := config.LoadInstagramConfig()
	// สำหรับใช้งาน api image
	image := image.NewImageHttp()

	// สำหรับการ Request ข้อมูลจาก API ภายนอก
	metaInstagramRequest := server.NewMetaInstagramRequest(igConfig.Api, igConfig.GraphVersion, igConfig.InstagramId, igConfig.AccessToken)
	// สำหรับใช้งาน API
	newMetaInstagramHandler := http.NewMetaInstagramHandler(metaInstagramRequest)

	// ส่งข้อมูลไปยัง Router เพื่อใช้งานฟังก์ชันการทำงานต่าง ๆ
	dependentcies := router.Dependencies{
		ImageHttpHandler:     image,
		MetaInstagramHandler: newMetaInstagramHandler,
	}

	// ใช้งาน Router
	r := router.Setup(&dependentcies)

	// โหลดข้อมูลสำหรับ APP เช่น port
	app := config.LoadAppConfig()
	// ตั้งค่าข้อมูลสำหรับ server
	serv := server.NewGinServer(app.AppAddress+app.AppPort, r)

	// สั่งให้ server เริ่มทำงาน
	serv.Start()

	// สำหรับ Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := serv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

}
