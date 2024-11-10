package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlPostData struct {
	Url string `json:"url" binding:"required"`
}

func GetUrl(c *gin.Context) {
	id := c.Param("id")
	// Perform Select call to the database

	found := true
	if found {
		c.JSON(http.StatusOK, gin.H{
			"msg": "helloooo " + id,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

func CreateUrl(c *gin.Context) {
	data := urlPostData{}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid data",
		})
	} else {
		// Perform Insert call to the database

		c.JSON(http.StatusCreated, gin.H{
			"msg": "CreateUrl",
		})
	}

}
