package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	cache *redis.Client
}

func GetRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("[ERR] Could not connect to Redis. Please make sure you have Redis Server running on localhost:6379.")
	}
	return &RedisClient{
		cache: rdb,
	}
}
