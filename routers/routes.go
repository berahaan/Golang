package routers

import (
	"GOLANG/auth"
	"GOLANG/handlers"
	"GOLANG/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", services.GetIntroduce)
	router.POST("/addusers", handlers.AddUsers)
	router.GET("/Alhuha", handlers.GetUsers)
	router.GET("/Alhuha/:id", handlers.GetUserById)
	router.POST("/signup", auth.HandleSignup)

}
