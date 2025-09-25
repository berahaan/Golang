package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"time"

	"gorm.io/gorm"
)

type RateLimitServices struct{}

func NewOTPService() *RateLimitServices {
	return &RateLimitServices{}
}

// CheckOTPAttempts checks if OTP has exceeded max attempts
func (rls *RateLimitServices) CheckOtpAttempts(userId uint, code string) (bool, string) {
	var otp models.OTP
	result := database.DB.Where("user_id=? AND code=?", userId, code).First(&otp)
	if result.Error != nil {
		return false, "OTP not found"
	}
	if time.Now().After(otp.ExpiresAt) {
		return false, "OTP has expired"
	}

	if otp.Attempts >= otp.MaxAttempts {
		return false, "Too many attempts . Please request a new OTP "
	}
	return true, ""
}

// IncrementOtpAttempts increment the attempts counter in database
func (rls *RateLimitServices) IncrementOtpAttempts(userId uint, code string) error {
	result := database.DB.Model(&models.OTP{}).Where("user_id=? AND code = ?", userId, code).Update("attempts", gorm.Expr("attempts+?", 1))
	return result.Error
}
