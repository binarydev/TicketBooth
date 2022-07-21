package datastore

import (
	"github.com/go-redis/redis/v8"
)

func OpenCacheConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
