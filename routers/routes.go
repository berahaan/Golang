package routers

import (
	"GOLANG/api/auth"
	"GOLANG/internals/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", services.GetIntroduce)
	router.POST("api/auth/signup", auth.HandleSignup)
	router.POST("api/auth/login", auth.HandleLoginAuth)
	router.POST("api/auth/verify-otp", auth.HandleVerifyOTP)
	router.POST("/api/auth/resend-otp", auth.HandleResendOtp)

}
