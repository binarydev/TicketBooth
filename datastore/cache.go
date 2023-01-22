package datastore

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var addr string
var pwd string
var db int

func InitializeRedis() *redis.Client {
	host := flag.String("redishost", "localhost", "Redis host address")
	port := flag.Int("redisport", 6379, "Redis port number")
	pwd = *flag.String("redispassword", "", "Redis password")
	db = *flag.Int("redisdbindex", 0, "Redis database index number")
	flag.Parse()
	addr = fmt.Sprintf("%s:%d", *host, *port)
	return OpenCacheConnection()
}

func OpenCacheConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
}
