package auth

import (
	"GOLANG/database"
	"GOLANG/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HandleSignup(c *gin.Context) {
	var input models.UserInput

	// Bind and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Check if user already exists
	var existingUser models.User
	result := database.DB.Where("email = ?", input.Email).First(&existingUser)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		// This handles actual database errors, not just "not found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User with this email already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process password",
		})
		return
	}

	// Create user
	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		// Add other fields from input if needed (e.g., Username, Name)
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user account",
		})
		return
	}

	// Return success response (consider returning the user ID or email)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": user.ID, // Useful for the client
		"email":   user.Email,
	})
}
