package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/api/controller"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createMongoClient() *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(os.Getenv("MONGO_DB")).Collection(collectionName)
}

func InsertTest(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Database(os.Getenv("MONGO_DB")).Collection("users").InsertOne(ctx, bson.D{{"hola", 1}})
}

func main() {
	mongoClient := createMongoClient()
	// InsertTest(mongoClient)
	fmt.Println(GetCollection(mongoClient, "users").FindOne(context.Background(), bson.D{{"hola", 1}}).DecodeBytes())
	defer func() {
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	r.GET("/:id", controller.GetUrl)
	r.POST("/", controller.CreateUrl)
	r.Run()

}
