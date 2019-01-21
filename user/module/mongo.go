package module

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var DB *mongo.Database

const url = "mongodb://127.0.0.1:27017"

func init() {
	client, err := mongo.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("connect mongo error: %v", err)
	}
	DB = client.Database("cinema_user")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("ping mongo error: %v", err)
	}
}
