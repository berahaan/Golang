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

func (rls *RateLimitServices) CheckOtpAttempts(userId uint, code string) (bool, string) {
	var otp models.OTP

	// FIX: Add code to the WHERE clause
	result := database.DB.Where("user_id = ? AND code = ? AND used = ?",
		userId, code, false).First(&otp)

	if result.Error != nil {
		log.Println("OTP not found or error:", result.Error)
		return false, "OTP not found"
	}

	log.Println("User informations in check otp attempts func Id, attempts, Max attempts",
		otp.UserID, otp.Attempts, otp.MaxAttempts)

	// Check if OTP is expired
	if time.Now().After(otp.ExpiresAt) {
		return false, "OTP has expired"
	}

	// Check if max attempts exceeded
	if otp.Attempts >= otp.MaxAttempts {
		return false, "Too many attempts. Please request a new OTP."
	}

	return true, ""
}

// IncrementOtpAttempts increment the attempts counter in database
func (rls *RateLimitServices) IncrementOtpAttempts(userId uint, code string) error {
	log.Println("Increment OTP attempts due to wrong code for user ", userId, "and code ", code)

	// FIX: Add code to the WHERE clause to target the specific OTP
	result := database.DB.Model(&models.OTP{}).
		Where("user_id = ? AND code = ? AND used = ?", userId, code, false).
		Update("attempts", gorm.Expr("attempts + ?", 1))

	if result.Error != nil {
		log.Println("Error incrementing attempts:", result.Error)
		return result.Error
	}

	switch result.RowsAffected {
	case 1:
		log.Println("✅ Successfully incremented attempts for user", userId, "code", code)
	case 0:
		log.Println("⚠️  No OTP found to increment - user", userId, "code", code, "might be expired/already used")
	default:
		log.Println("❌ Unexpected: Updated", result.RowsAffected, "rows for user", userId)
	}

	return nil
}
