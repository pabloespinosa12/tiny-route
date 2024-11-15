package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/api/middleware"
	"github.com/pabloespinosa12/tiny-route/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUrl(c *gin.Context) {
	db := middleware.GetDatabaseFromContext(c)
	id := c.Param("id")

	urlInfo := models.Url{}
	urlId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = db.GetCollection("urls").FindOne(c, bson.M{"_id": urlId}).Decode(&urlInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	found := true
	if found {
		c.Redirect(http.StatusPermanentRedirect, urlInfo.Url)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

func CreateUrl(c *gin.Context) {
	db := middleware.GetDatabaseFromContext(c)
	data := models.Url{}

	// TODO: This doesn't fail the request if the body is empty
	// TODO: Should check content type header
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid data",
		})
		return
	}

	// TODO: Post data validation

	// Perform Insert call to the database
	res, err := db.InsertOne("urls", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Internal server error",
		})
		return
	}

	mongoId := res.InsertedID.(primitive.ObjectID).Hex()
	c.JSON(http.StatusCreated, gin.H{
		"msg": mongoId,
	})

}
