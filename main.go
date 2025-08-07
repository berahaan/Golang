package main

import (
	"GOLANG/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// backend Journey Started now ....
	server := gin.Default()
	routers.RegisterRoutes(server)
	fmt.Println("Server running on port 8080")
	server.Run(":8080")

}
