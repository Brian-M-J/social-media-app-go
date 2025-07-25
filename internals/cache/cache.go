package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Client() *redis.Client {
	return cache
}

func Connect() {
	ctx := context.Background()

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "localhost:6379"
		cache = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	} else {
		redisOptions, err := redis.ParseURL(redisUrl)
		if err != nil {
			log.Fatalln("Failed to parse URL:", err)
		}
		cache = redis.NewClient(redisOptions)
	}

	cmd := cache.Ping(ctx)
	if cmd.Err() != nil {
		log.Fatalf("Error connecting caching database: %v\n", cmd.Err())
	}

	fmt.Printf("Successfully connected to the Redis cache")

}
