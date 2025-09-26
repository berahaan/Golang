package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"time"
)

func StoreOTP(User_Id uint, otpNumber string, expTime time.Time) error {
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

	// 1. Check Rate Limit and Get Active OTP Record
	otpRecord, errMsg := ratelimitService.CheckOtpAttempts(userId)

	if otpRecord == nil {
		// No active OTP found for the user
		return false, "Invalid or expired OTP. " + errMsg
	}

	// 2. Handle Attempts/Expiration Exhausted (from CheckOtpAttempts)
	if errMsg != "" {
		return false, errMsg
	}

	// 3. Check if the entered code is WRONG
	if otpRecord.Code != code {
		// The code is WRONG. Increment the attempt counter for the *valid* OTP record.
		ratelimitService.IncrementOtpAttempts(userId, otpRecord.Code)

		// Re-check the attempt count after incrementing
		if otpRecord.Attempts+1 >= otpRecord.MaxAttempts {
			return false, "Too many attempts. Please request a new OTP"
		}

		return false, "Invalid Security code."
	}

	// 4. ATOMIC VERIFICATION & USAGE (If code is correct and passed initial checks)
	result := database.DB.Model(&models.OTP{}).
		Where("user_id = ? AND code = ? AND used = ? AND expires_at > ?",
			userId, code, false, time.Now()).
		Update("used", true)

	if result.RowsAffected == 1 {
		// SUCCESS: OTP consumed. Do not increment attempts.
		return true, ""
	}

	return false, "Verification Failed."
}
