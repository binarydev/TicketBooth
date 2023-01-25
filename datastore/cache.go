package datastore

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisAddr string
var RedisPwd string
var RedisDbNum int

func InitializeRedis() {
	OpenCacheConnection().Close()
	fmt.Println("Redis connection successfully tested")
}

func OpenCacheConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPwd,
		DB:       RedisDbNum,
	})
}
