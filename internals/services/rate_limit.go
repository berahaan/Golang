package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type RateLimitServices struct{}

func NewOTPService() *RateLimitServices {
	return &RateLimitServices{}
}

func (rls *RateLimitServices) CheckOtpAttempts(userId uint) (*models.OTP, string) {
	var otp models.OTP
	result := database.DB.Where("user_id = ?", userId).First(&otp)

	log.Println("\n User informations in check otp attempts func Id, attempts,Max attempts \n\n ", otp.ID, otp.Attempts, otp.MaxAttempts)
	if result.Error != nil {
		// No active OTP found for this user
		return nil, "Please request a new OTP."
	}

	if otp.Attempts >= otp.MaxAttempts {
		log.Println("otp.Attempts >= Otp.MaxAttempts ", otp.Attempts, otp.MaxAttempts)
		return &otp, "Too many attempts. Please request a new OTP" // Return the record with the error
	}

	if time.Now().After(otp.ExpiresAt) {
		return &otp, "OTP has expired" // Return the record with the error
	}

	// OTP exists, is not expired, and has attempts remaining
	return &otp, ""
}

// IncrementOtpAttempts increment the attempts counter in database
func (rls *RateLimitServices) IncrementOtpAttempts(userId uint, code string) error {
	log.Println("Increment OTP atttempts due to wrong code for user ", userId, "and code ", code)
	result := database.DB.Model(&models.OTP{}).Where("user_id=? AND code = ?", userId, code).Update("attempts", gorm.Expr("attempts+?", 1))
	if result.RowsAffected == 1 {
		log.Println("Rows affected = 1 hitted")
	}
	return result.Error
}
