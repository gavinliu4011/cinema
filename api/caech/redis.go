package caech

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

var Redis *redis.Client

func init() {
	redisAddr := os.Getenv("REDIS_ADDRESS")
	if redisAddr == "" {
		redisAddr = "local:6379"
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatal("redis conn error:", err)
	}
}
