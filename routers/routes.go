package routers

import (
	// "GOLANG/handlers"
	"GOLANG/handlers"
	"GOLANG/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", services.GetIntroduce)
	router.POST("/addusers", handlers.AddUsers)

}
