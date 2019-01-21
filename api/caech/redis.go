package caech

import (
	"github.com/go-redis/redis"
	"log"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatal("redis conn error:", err)
	}
}
