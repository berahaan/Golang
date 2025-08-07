package routers

import (
	"GOLANG/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", services.GetIntroduce)
	router.GET("/albums", services.GetAlbums)
	router.GET("/albums/:id", services.GetAlbumByID)
}
