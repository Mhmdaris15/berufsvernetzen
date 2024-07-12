package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// client *mongo.Client
	ctx            context.Context
	UserCollection = GetCollection(DB, "Users")
)

type MongoDBService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBService() (*MongoDBService, error) {
	clientOptions := options.Client().ApplyURI(configs.EnvMongoURI())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	collection := client.Database("berufsvernetzen").Collection("Items")

	return &MongoDBService{
		client:     client,
		collection: collection,
	}, nil
}

func ConnectDB() *mongo.Client {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configs.EnvMongoURI()).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Create Collecion for Users with primary key id
	// UserCollection := client.Database("berufsvernetzen").Collection("Users")
	// indexModel := mongo.IndexModel{
	// 	Keys:    bson.M{"id": 1},
	// 	Options: options.Index().SetUnique(true),
	// }

	// _, err = UserCollection.Indexes().CreateOne(ctx, indexModel)

	return client
}

var DB *mongo.Client = ConnectDB()

func DisconnectDB(client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(configs.EnvDatabaseName()).Collection(collectionName)
	return collection
}
