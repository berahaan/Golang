package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleVerifyOTP(c *gin.Context) {
	log.Println("Verify OTP Endpoint Hit")
	var input models.VerifyOTPInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind json",
		})
		return
	}
	if allowed, messg := services.VerifyOTP(input.UserId, input.Code); !allowed {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or Expired OTP" + messg,
		})
		return
	}
	var user models.User
	if err := database.DB.First(&user, input.UserId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
		return
	}
	token, err := services.GenerateJWTtoken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": " Failed to generate JWT token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})

}
