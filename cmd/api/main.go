package main

import (
	"api-for-shops-on-instagram/config"
	server "api-for-shops-on-instagram/internal/infrastructure/server/http"
	image "api-for-shops-on-instagram/internal/module/image/delivery/http"
	http "api-for-shops-on-instagram/internal/module/instagram/delivery/http"
	"api-for-shops-on-instagram/internal/module/instagram/usecase"
	"api-for-shops-on-instagram/internal/router"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	igConfig := config.LoadInstagramConfig()
	igUc := usecase.NewInstagramInfoUsecase(igConfig.Api, igConfig.GraphVersion, igConfig.AccessToken)
	igHandler := http.NewInstagramHandler(igUc)
	image := image.NewImageHttp()

	dependentcies := router.Dependencies{
		InstagramHandler: igHandler,
		ImageHttpHandler: image,
	}

	r := router.Setup(&dependentcies)

	app := config.LoadAppConfig()
	serv := server.NewGinServer(app.AppAddress+app.AppPort, r)

	serv.Start()

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
