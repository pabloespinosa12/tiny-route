package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/api/middleware"
)

type urlPostData struct {
	Url string `json:"url" binding:"required"`
}

func GetUrl(c *gin.Context) {
	db := middleware.GetDatabaseFromContext(c)
	id := c.Param("id")

	// Perform Select call to the database
	db.PingTest() // Avoid unused variable error

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
	db := middleware.GetDatabaseFromContext(c)
	data := urlPostData{}

	db.PingTest() // Avoid unused variable error

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
