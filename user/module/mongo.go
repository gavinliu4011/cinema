package module

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"os"
)

var DB *mongo.Database

func init() {
	mongodbAddr := os.Getenv("MONGO_ADDRESS")
	if mongodbAddr == "" {
		mongodbAddr = "mongodb://127.0.0.1:27017"
	}
	client, err := mongo.Connect(context.Background(), mongodbAddr)
	if err != nil {
		log.Fatalf("connect mongo error: %v", err)
	}
	DB = client.Database("cinema_user")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("ping mongo error: %v", err)
	}
}
