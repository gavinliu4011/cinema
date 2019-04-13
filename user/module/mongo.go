package module

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	mongodbAddr := os.Getenv("MONGO_ADDRESS")
	if mongodbAddr == "" {
		mongodbAddr = "mongodb://127.0.0.1:27017"
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbAddr))
	if err != nil {
		log.Fatalf("connect mongo error: %v", err)
	}
	DB = client.Database("cinema_user")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("ping mongo error: %v", err)
	}
}
