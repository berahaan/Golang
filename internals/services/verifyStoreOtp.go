package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"fmt"
	"log"
	"time"
)

func StoreOTP(User_Id uint, otpNumber string, expTime time.Time) error {
	checkAttempts := NewOTPService()
	log.Println("Store OTP")
	if allowed, message := checkAttempts.CheckOtpAttempts(User_Id, otpNumber); !allowed {
		return fmt.Errorf("Rate limit exceeded: %s", message)
	}

	otp := models.OTP{
		UserID:      User_Id,
		Code:        otpNumber,
		ExpiresAt:   expTime,
		Used:        false,
		Attempts:    0,
		MaxAttempts: 3,
	}

	//  now insert this information to db OTP
	return database.DB.Create(&otp).Error

}

func VerifyOTP(userId uint, code string) (bool, string) {
	ratelimitService := NewOTPService()

	// 1. Check if OTP is within the allowed attempts (Pre-check for exhaustion)
	if allowed, mess := ratelimitService.CheckOtpAttempts(userId, code); !allowed {
		return false, mess
	}

	// 2. ATOMIC VERIFICATION & USAGE (Single Database Command)
	// Attempt to find a match AND update the 'used' status in ONE step.
	result := database.DB.Model(&models.OTP{}).
		Where("user_id = ? AND code = ? AND used = ? AND expires_at > ?",
			userId, code, false, time.Now()).
		Update("used", true)

	// 3. Check for Database Errors
	// +97470908074
	if result.Error != nil {
		// A database connection/system error occurred.
		return false, "Verification Failed due to a system error."
	}

	// 4. Check for Success
	if result.RowsAffected == 1 {
		// An unused, unexpired, and matching OTP was atomically consumed.
		// DO NOT INCREMENT attempts here, as the attempt was successful.
		return true, ""
	}

	// 5. Handle Failure (RowsAffected == 0)
	// No row was updated, meaning one of the WHERE conditions failed (wrong code, already used, or expired).

	// Increment the attempts counter because the verification failed.
	ratelimitService.IncrementOtpAttempts(userId, code)

	// Provide a better error message (optional, but good UX)
	var otp models.OTP
	// Read the current state to determine the specific failure reason
	if err := database.DB.Where("user_id = ? AND code = ?", userId, code).First(&otp).Error; err == nil {
		if otp.Used {
			return false, "OTP already used."
		}
		if time.Now().After(otp.ExpiresAt) {
			return false, "OTP has expired."
		}
	}

	// Default/catch-all error message
	return false, "Invalid Security code."
}
