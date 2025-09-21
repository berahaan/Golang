package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	"GOLANG/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLoginAuth(c *gin.Context) {
	// now we need to take emails and password of a users from clients
	var LoginInput models.LoginInput

	if err := c.ShouldBindJSON(&LoginInput); err != nil {
		c.JSON(400, gin.H{
			"error": "unable to bind json" + err.Error(),
		})
		return
	}
	//  sanitize the email and password
	LoginInput.Email = utils.SanitizeEmail(LoginInput.Email)
	LoginInput.Password = utils.SanitizePassword(LoginInput.Password)

	if !utils.ValidateEmail(LoginInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Emails Please provide a valid emails",
		})
		return
	}
	// find user by emails
	var user models.User
	result := database.DB.Where("email=?", LoginInput.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password ",
		})
		return
	}
	// compare password with stored hash password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password",
		})
	}
	// generate JWT with claims
	token, err := services.GenerateJWTtoken(user.Email, int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": " Failed to generate JWT token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"Token":   token,
		"User": gin.H{
			"id":    user.ID,
			"Email": user.Email,
		},
	})

}
