package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	emails "GOLANG/pkg/email"
	"GOLANG/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleResetPassWord(c *gin.Context) {
	// marsh the encoming emails of users
	var resetEmails models.ResetPasswordInput
	if err := c.ShouldBindJSON(&resetEmails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind json",
		})
		return
	}
	// we need to check if user with specified emails is exist in database or not
	var user models.User
	resetEmails.Emails = utils.SanitizeEmail(resetEmails.Emails)
	if !utils.ValidateEmail(resetEmails.Emails) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid email address",
		})
		return
	}

	result := database.DB.Where("emails=?", resetEmails.Emails).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "If this email exists, a reset code has been sent ",
		})
		return
	}
	// user exist so now we need to send the a otp code
	otpCode := utils.GenerateOTP(6)
	expiresAt := time.Now().Add(5 * time.Minute)
	// Store OTP in database
	storeOtpResult := services.StoreOTP(user.ID, otpCode, expiresAt)
	if storeOtpResult != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to store OTP",
		})
		return
	}
	subject := "Your Login OTP Code"
	body := "Hello, your OTP code for Password reset is: " + otpCode + "\nIt will expire in 5 minutes. \n Please don't share with anyone"
	// Send OTP Emails
	if err := emails.SendEmailsTo(user.Email, subject, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send OTP email" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "OTP sent to your emails please check your emails",
		"userID":  user.ID,
	})

}
