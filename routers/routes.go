package routers

import (
	"GOLANG/handlers"
	"GOLANG/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", services.GetIntroduce)
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/addAlbum", handlers.PostAlbums)
	router.DELETE("/removeAlbums/:id", handlers.RemoveAlbumByID)
}
