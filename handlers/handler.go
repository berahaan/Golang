package handlers

import (
	"GOLANG/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Album One", Artist: "Artist one ", Price: 8.66},
	{ID: "2", Title: "Album two", Artist: "Artist one ", Price: 5.66},
	{ID: "3", Title: "Album Three", Artist: "Artist one ", Price: 2.66},
}

// let us proceed with delete the specific albums from the slices

func RemoveAlbumByID(c *gin.Context) {
	// let us track the id of the albums
	id := c.Param("id")
	found := false
	// looping through albums
	for index, album := range albums {
		if album.ID == id {
			// now we got the id of albums to be deleted
			albums = append(albums[:index], albums[index+1:]...)
			found = true
			break
		}
	}
	if found {
		c.JSON(200, gin.H{
			"Message": "successfully removed",
		})
		return
	} else {
		c.JSON(404, gin.H{
			"message": "Albums to be deleted is not found",
		})
	}

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
	// we need to handle if the albums requested is not exist
	c.JSON(404, gin.H{
		"Message": "Requested albums not found sorry ",
	})
}

// fetching the all albums here
func GetAlbums(c *gin.Context) {
	fmt.Println("GetAlbums called ")
	c.JSON(200, albums)
}

// introduce any one thatb access my routes explicitly
func GetIntroduce(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello this is from Port 8080 local host server",
	})
}

// updates the albums

func UpdateAlbum(c *gin.Context) {
	// let us extract id from requests
	var UpdateAlbums models.Album
	id := c.Param("id")
	if err := c.BindJSON(&UpdateAlbums); err != nil {
		c.JSON(404, gin.H{
			"Message": "Error while serializing Json",
		})
	}
	for index, album := range albums {
		if album.ID == id {
			// we need to update it
			albums[index].ID = UpdateAlbums.ID
			albums[index].Artist = UpdateAlbums.Artist
			albums[index].Price = UpdateAlbums.Price
			albums[index].Title = UpdateAlbums.Title
			c.JSON(200, gin.H{
				"message": "Successfully added",
				"updated": albums[index],
			})
		}
	}
	c.JSON(404, gin.H{
		"message": "Albums to be Updated not found now ...",
	})
}

// now let us build the post request to add the albums for clients
func PostAlbums(c *gin.Context) {
	// now let us add some albums to the existing albums
	var newAlbums models.Album
	if err := c.ShouldBindJSON(&newAlbums); err != nil {
		// now this mean if the json is not destructured to this albums struct we need to send errors
		c.JSON(404, gin.H{
			"error": "cannot we destructure the Jsons here ",
		})
	}
	// now adding the encoming json data to here
	albums = append(albums, newAlbums)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Successfully added new albums",
		"Albums":  newAlbums,
	})

}
