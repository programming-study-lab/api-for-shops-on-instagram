package router

import (
	"net/http"
	"time"

	image "api-for-shops-on-instagram/internal/module/image/delivery/http"
	instagram "api-for-shops-on-instagram/internal/module/instagram/delivery/http"
	v1 "api-for-shops-on-instagram/internal/router/v1"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	InstagramHandler *instagram.InstagramHandler
	ImageHttpHandler *image.ImageHttpHandler
}

func Setup(dependencies *Dependencies) *gin.Engine {
	// func Setup(router *gin.Engine) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(gin.Recovery())

	v1Group := router.Group("/api/v1")
	{
		v1Group.GET("/test", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(
				http.StatusOK,
				gin.H{
					"status":  true,
					"message": "success",
					"data":    time.Now().Format(time.DateTime),
				},
			)
		})
		v1.ImageRouter(v1Group, dependencies.ImageHttpHandler)
		v1.InstagramRouter(v1Group, dependencies.InstagramHandler)
	}

	return router
}
