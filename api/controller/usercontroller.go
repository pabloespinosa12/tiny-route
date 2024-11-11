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

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	dbService := middleware.GetDatabaseFromContext(c)
	err = dbService.GetCollection("users").FindOne(c, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
