package services

import (
	"GOLANG/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// let us creates a struct to represent the encoming Albums data

// let us send some albums when a client is requesitons for albums here
var albums = []models.Album{
	{ID: "1", Title: "Album One", Artist: "Artist one ", Price: 8.66},
	{ID: "2", Title: "Album two", Artist: "Artist one ", Price: 5.66},
	{ID: "3", Title: "Album Three", Artist: "Artist one ", Price: 2.66},
}

func GetAlbums(c *gin.Context) {
	fmt.Println("GetAlbums called ")
	c.JSON(200, albums)
}
func GetIntroduce(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello this is from Port 8080 local host server",
	})
}

func GetAlbumByID(c *gin.Context) {
	// so we need to get the id of albums clients is requesting for
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.JSON(200, album)
			return
		}
	}
	c.JSON(404, gin.H{
		"Message": "Requested albums not found sorry ",
	})
}
