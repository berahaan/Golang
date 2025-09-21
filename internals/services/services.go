package services

import (
	"github.com/gin-gonic/gin"
)

// let us send some albums when a client is requesitons for albums here

func GetIntroduce(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello this is from Port 8080 local host server",
	})
}
