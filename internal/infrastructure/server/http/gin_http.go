package http

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	HttpServer *http.Server
}

func NewGinServer(address string, router *gin.Engine) *GinServer {

	// if router != nil {
	// 	log.Println("ไม่พบ router")
	// }

	return &GinServer{
		HttpServer: &http.Server{
			Addr:    address,
			Handler: router,
			// ReadTimeout:                  10 * time.Second,
			// WriteTimeout:                 10 * time.Second,
			// DisableGeneralOptionsHandler: true,
			// IdleTimeout:                  120 * time.Second,
		},
	}
}

func (g *GinServer) Start() {
	go func() {
		slog.Info("Gin Server is listening", "address", g.HttpServer.Addr)

		log.Printf("Server is running...")

		err := g.HttpServer.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Gin server error: ", err.Error())
			os.Exit(1)
		}

		// if err != nil && !errors.Is(err, http.ErrServerClosed) {
		// 	slog.Error("Gin server crashed unexpectedly", "error", err)
		// 	os.Exit(1)
		// }

	}()
}

func (g *GinServer) Shutdown(ctx context.Context) error {

	slog.Info("Gin Server is rejecting new requests and dringing current connections...")

	return g.HttpServer.Shutdown(ctx)
}
