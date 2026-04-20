package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("[ERR] Failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(opts)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("[ERR] Failed to connect to Redis: %v", err)
	}

	log.Printf("[INFO] Connected to Redis")
	return client
}
