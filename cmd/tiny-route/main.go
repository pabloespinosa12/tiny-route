package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/api/controller"
	"github.com/pabloespinosa12/tiny-route/internal/database"
)

func main() {
	database := database.NewDatabaseService(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	defer database.CloseConnection()

	// TODO: Create middleware that adds the database service to the Gin context

	r := gin.Default()
	r.GET("/:id", controller.GetUrl)
	r.POST("/", controller.CreateUrl)
	r.Run()

}
