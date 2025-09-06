package main

// just write down a folder name and then follow by package name
import (
	"GOLANG/database"
	"GOLANG/models"
	"GOLANG/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Album{})
	server := gin.Default()
	routers.RegisterRoutes(server)
	fmt.Println("Server running on port 8080")
	server.Run(":8080")
}
