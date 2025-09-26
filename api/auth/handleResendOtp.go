package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	emails "GOLANG/pkg/email"
	"GOLANG/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResendOTPInput struct {
	UserID uint `json:"user_id" `
}

func HandleResendOtp(c *gin.Context) {
	var input ResendOTPInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid JSON",
		})
	}
	result := database.DB.Model(&models.OTP{}).Where("user_id", input.UserID).Update("attempts", 0)
	if result.Error != nil {
		log.Println("result.Error in HandleResend OTP ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update the database column attempts",
		})
		return
	}
	// Get user details like emails to send otp codes
	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found ",
		})
		return
	}
	// generate new OTP
	num := utils.GenerateOTP(6)
	otp_expires := time.Now().Add(5 * time.Minute)
	// Store OTP in database along with User ID
	err := services.StoreOTP(user.ID, num, otp_expires)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to store OTP",
		})
		return
	}
	subject := "Your Login OTP Code"
	body := "Hello, your OTP code is: " + num + "\nIt will expire in 5 minutes."
	// Send OTP Emails
	if err := emails.SendEmailsTo(user.Email, subject, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send OTP email" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": " OTP Resent to your emails please check your emails",
		"UserId":  user.ID,
	})

}
