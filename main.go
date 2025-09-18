package main

// just write down a folder name and then follow by package name for importing purposes
import (
	"GOLANG/database"
	"GOLANG/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
		return
	}

	database.Connect()
	database.MigrateTables()
	server := gin.Default()
	routers.RegisterRoutes(server)
	fmt.Println("Server running on port 8080")
	server.Run(":8080")
}
