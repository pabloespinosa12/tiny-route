package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/api/controller"
	"github.com/pabloespinosa12/tiny-route/api/middleware"
	"github.com/pabloespinosa12/tiny-route/internal/database"
)

func main() {
	database := database.NewDatabaseService(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	defer database.CloseConnection()

	r := gin.Default()
	r.Use(middleware.DatabaseProvider(database))
	r.GET("/:id", controller.GetUrl)
	r.POST("/", controller.CreateUrl)
	r.GET("/user/:id", controller.GetUserById)
	r.Run()
}
