package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// we need to take the user code and verify if exist
type VerifyEmail struct {
	UserID uint   `json:"user_id"`
	Code   string `json:"code"`
}

func HandleVerifyUserEmails(c *gin.Context) {
	var verfyEmail VerifyEmail
	if err := c.ShouldBindJSON(&verfyEmail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error marshaling Json",
		})
		return
	}
	var otpRecord models.OTP
	result := database.DB.Where("user_id = ? AND code = ? AND used = ? AND expires_at > ?",
		verfyEmail.UserID, verfyEmail.Code, false, time.Now()).First(&otpRecord)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or expired verification code",
		})
		return
	}

	if otpRecord.Code == verfyEmail.Code {
		if time.Now().After(otpRecord.ExpiresAt) {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "OTP Expired",
			})
			return
		}
		database.DB.Model(&otpRecord).Update("used", true)
		//  otp is valid
		c.JSON(http.StatusAccepted, gin.H{
			"Message": "OTP verified Successfully",
		})
	}
	//   otp is invalid her
	c.JSON(http.StatusUnauthorized, gin.H{
		"error":  "OTP is invalid",
		"UserID": verfyEmail.UserID,
	})

}
