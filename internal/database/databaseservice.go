package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService struct {
	client   *mongo.Client
	database string
}

func NewDatabaseService(connectionString string, database string) *DatabaseService {
	client := createMongoClient(connectionString)

	dbService := &DatabaseService{
		client:   client,
		database: database,
	}

	if !dbService.PingTest() {
		panic(fmt.Sprintf("Ping to database failed"))
	}

	return dbService
}

func createMongoClient(connectionString string) *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

func (db *DatabaseService) PingTest() bool {
	var result bson.M
	if err := db.client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return false
	}
	return true
}

// TODO: Extend available methods to support all CRUD operations
func (db *DatabaseService) GetCollection(collectionName string) *mongo.Collection {
	return db.client.Database(db.database).Collection(collectionName)
}

func (db *DatabaseService) InsertTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.client.Database(db.database).Collection("users").InsertOne(ctx, bson.D{{"hola", 1}})
}

func (db *DatabaseService) CloseConnection() {
	err := db.client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
