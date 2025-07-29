package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// backend Journey Started now ....
	server := gin.Default()
	fmt.Println("Server running on port 8080")
	server.Run(":8080")
}
