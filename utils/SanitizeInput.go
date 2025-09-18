package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/microcosm-cc/bluemonday"
)

func ValidateEmail(email string) bool {
	// Basic email regex pattern
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
func SanitizeString(input string) string {
	// Create a strict policy that only allows plain text
	policy := bluemonday.StrictPolicy()
	sanitized := policy.Sanitize(input)

	// Additional basic sanitization: trim spaces
	sanitized = strings.TrimSpace(sanitized)

	return sanitized
}

// SanitizeEmail specifically sanitizes email input
func SanitizeEmail(email string) string {
	email = strings.TrimSpace(email)
	email = SanitizeString(email)
	return email
}
func ValidatePaswordStrength(password string) (bool, string) {
	// we need to check if the passwords have the necessary qualification or the client need to satify the necessary conditions here
	fmt.Println("Validate passwords ")
	var hasUpper, hasLower, hasSpecialChar, hasNumber bool
	if len(password) < 5 {
		return false, "Password need to be greater than 5 characters "
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}
	if !hasUpper {
		return false, "Password must have at least one Uppercase letter"
	}
	if !hasLower {
		return false, "Password must have at least one Lowercase letter"
	}
	if !hasNumber {
		return false, "password need to have at least one Numbers"
	}
	if !hasSpecialChar {
		return false, "Password need to have at least 1 special characters"
	}

	return true, ""

}
func SanitizePassword(password string) string {
	// Remove whitespace (including newlines, tabs)
	password = strings.TrimSpace(password)

	// Create a custom policy for passwords
	policy := bluemonday.NewPolicy()
	// We don't allow any HTML tags in passwords
	policy.AllowElements("") // No elements allowed
	// Remove common SQL injection patterns (very basic example)
	dangerousPatterns := []string{
		"'", "\"", "--", "/*", "*/", ";", "=",
		"<script", "</script>", "javascript:",
		"onload=", "onerror=", "onclick=",
	}

	for _, pattern := range dangerousPatterns {
		password = strings.ReplaceAll(password, pattern, "")
	}

	return password
}
