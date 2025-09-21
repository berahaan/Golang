package auth

import (
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleVerifyOTP(c *gin.Context) {
	var input models.VerifyOTPInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind json",
		})
		return
	}
	if !services.VerifyOTP(input.UserId, input.Code) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or Expired OTP",
		})
		return
	}
	token, err := services.GenerateJWTtoken(int(input.UserId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": " Failed to generate JWT token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "OTP Verified Successfully",
		"token":   token,
	})

}
