package datastore

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func OpenCacheConnection() *redis.Client {
	host := flag.String("redishost", "localhost", "Redis host address")
	pwd := flag.String("redispassword", "", "Redis password")
	port := flag.Int("redisport", 6379, "Redis port number")
	dbnum := flag.Int("redisdbindex", 0, "Redis database index number")
	flag.Parse()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", *host, *port),
		Password: *pwd,
		DB:       *dbnum,
	})
}
