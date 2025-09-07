package handlers

import (
	"GOLANG/database"
	"GOLANG/models"
	"log"

	"github.com/gin-gonic/gin"
)

// now let us create a functions that will implement this stuct and parse jsons
func AddUsers(c *gin.Context) {
	// we need to create a variable of Type Allhuha to parse the incoming json
	var NewUser models.Allhuha
	if err := c.BindJSON((&NewUser)); err != nil {
		c.JSON(404, gin.H{
			"Error": err.Error(),
		})
		return
	}
	// save the users to database
	result := database.DB.Create(&NewUser)
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		c.JSON(500, gin.H{
			"Error": "Error while saving to database",
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Now rows affected",
		})
		return
	}
	// now we need to send the acknowledgement to the users or clients
	c.JSON(201, gin.H{
		"Message": " User Added Successfully",
		"Users":   NewUser,
	})
}
