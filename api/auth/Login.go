package auth

import (
	"GOLANG/internals/database"
	"GOLANG/internals/models"
	"GOLANG/internals/services"
	"GOLANG/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLoginAuth(c *gin.Context) {
	// now we need to take emails and password of a users from clients
	fmt.Println("Login Endpoint Hit")
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
	// find user by emails and get all information to the user variables
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
		return
	}
	// Create a functions to generate and send OTP of 6 digit
	num := utils.GenerateOTP(6)
	otp_expires := time.Now().Add(5 * time.Minute)
	// Store OTP in database along with User ID
	err := services.StoreOTP(user.ID, num, otp_expires)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to store OTP",
		})
		return
	}
	// Send OTP Emails
	// generate JWT with claims
	token, err := services.GenerateJWTtoken(int(user.ID))
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
