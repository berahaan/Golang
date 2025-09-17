package utils

import (
	"regexp"
	"strings"

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
