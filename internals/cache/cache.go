package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Client() *redis.Client {
	return cache
}

func Connect() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cmd := rdb.Ping(ctx)
	if cmd.Err() != nil {
		log.Fatalf("Error connecting caching database: %v\n", cmd.Err())
	}

	fmt.Printf("Successfully connected to the Redis cache")
	cache = rdb

}
