package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/binarydev/ticketbooth/datastore"
	"github.com/binarydev/ticketbooth/router"
)

func main() {
	log.SetFlags(log.Ldate | log.LUTC)
	initFlags()
	datastore.InitializeDatabase()
	router.Engine.Run(":8080")
}

func initFlags() {
	dbhost := flag.String("dbhost", "localhost", "Database host address")
	dbuser := flag.String("dbuser", "postgres", "Database username")
	dbpwd := flag.String("dbpassword", "dev", "Database password")
	dbname := flag.String("dbname", "dev", "Database schema name")
	dbport := flag.Int("dbport", 5432, "Database port number")
	skipDbMigrations := flag.Bool("skip-db-migrations", false, "Skip database schema migrations")

	redishost := flag.String("redishost", "localhost", "Redis host address")
	redisport := flag.Int("redisport", 6379, "Redis port number")
	redispwd := flag.String("redispassword", "", "Redis password")
	redisDbIndex := flag.Int("redisdbindex", 0, "Redis database index number")

	flag.Parse()

	datastore.SkipMigrateDB = skipDbMigrations
	datastore.RedisPwd = *redispwd
	datastore.RedisDbNum = *redisDbIndex
	datastore.RedisAddr = fmt.Sprintf("%s:%d", *redishost, *redisport)
	datastore.PostgresDsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/New_York",
		*dbhost,
		*dbuser,
		*dbpwd,
		*dbname,
		*dbport,
	)
}
