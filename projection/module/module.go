package module

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

const url = "mongodb://127.0.0.1:27017"

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("connect mongo error: %v", err)
	}
	DB = client.Database("cinema_projection")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("ping mongo error: %v", err)
	}
}
