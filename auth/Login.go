package auth

import (
	"GOLANG/database"
	"GOLANG/models"
	"GOLANG/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleLoginAuth(c *gin.Context) {
	// now we need to take emails and password of a users from clients
	var LoginInput models.LoginInput
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

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
	result := database.DB.Where("email=?", LoginInput.Email).First(&models.User{})
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password ",
		})
		return
	}

}
