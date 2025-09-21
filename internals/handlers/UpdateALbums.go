package handlers

import (
	"GOLANG/internals/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdateAlbums(c *gin.Context) {
	fmt.Println("Update Albums called ")
	Id := c.Param("id")
	var updateBody models.Album
	body := c.BindJSON(&updateBody)
	fmt.Println(body, Id)
	// now we need to find a user with specified ID then we need to update his informations as well

}
