package config

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

func NewRedisClient() *redis.Client {
	redisHost := Env("REDIS_HOST", "127.0.0.1")
	redisPort := Env("REDIS_PORT", "6379")
	redisPass := Env("REDIS_PASSWORD", "redisPass")

	// format string
	redisURL := fmt.Sprintf("%s:%s", redisHost, redisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0,
	})

	return client
}
