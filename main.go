package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// backend Journey Started now ....
	server := gin.Default()

	fmt.Println("Server running on port 8080")
	server.GET("/", func(c *gin.Context) {
		// when this routes is called send the Json data called Message to the clients
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	server.Run(":8080")

}
