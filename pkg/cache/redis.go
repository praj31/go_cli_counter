package cache

import (
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
	return &RedisClient{
		cache: rdb,
	}
}
