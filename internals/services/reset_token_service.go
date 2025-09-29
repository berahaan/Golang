package services

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type ResetTokenService struct{}

func NewResetTokenService() *ResetTokenService {
	return &ResetTokenService{}
}

// GenerateResetToken creates a secure random token for password reset
func (rts *ResetTokenService) GenerateResetToken(userID uint) (string, error) {
	// Generate secure random token
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	token := hex.EncodeToString(tokenBytes)

	// Set expiration (1 hour from now)
	expiresAt := time.Now().Add(1 * time.Hour)

	// Store token in database
	resetToken := models.PasswordResetToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		Used:      false,
	}

	if err := database.DB.Create(&resetToken).Error; err != nil {
		return "", fmt.Errorf("failed to store reset token: %w", err)
	}

	return token, nil
}

// ValidateResetToken checks if a reset token is valid and not expired/used
func (rts *ResetTokenService) ValidateResetToken(token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken

	result := database.DB.Where("token = ? AND used = ? AND expires_at > ?",
		token, false, time.Now()).First(&resetToken)

	if result.Error != nil {
		return nil, fmt.Errorf("invalid or expired reset token")
	}

	return &resetToken, nil
}

// MarkTokenAsUsed invalidates a reset token after use
func (rts *ResetTokenService) MarkTokenAsUsed(token string) error {
	result := database.DB.Model(&models.PasswordResetToken{}).
		Where("token = ?", token).
		Update("used", true)

	if result.Error != nil {
		return fmt.Errorf("failed to mark token as used: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("token not found")
	}

	return nil
}

// CleanupExpiredTokens removes old expired tokens (can be run as cron job)
func (rts *ResetTokenService) CleanupExpiredTokens() error {
	result := database.DB.Where("expires_at < ? OR used = ?", time.Now(), true).Delete(&models.PasswordResetToken{})

	if result.Error != nil {
		return fmt.Errorf("failed to cleanup expired tokens: %w", result.Error)
	}

	fmt.Printf("Cleaned up %d expired reset tokens\n", result.RowsAffected)
	return nil
}
