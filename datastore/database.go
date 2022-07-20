package datastore

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn string

func InitializeDatabase() *gorm.DB {
	shouldMigrateDB := flag.Bool("skip-db-migrations", true, "Skip database schema migrations")
	dbhost := flag.String("dbhost", "localhost", "Database host address")
	dbuser := flag.String("dbuser", "postgres", "Database username")
	dbpwd := flag.String("dbpassword", "dev", "Database password")
	dbname := flag.String("dbname", "dev", "Database schema name")
	dbport := flag.Int("dbport", 5432, "Database port number")
	flag.Parse()
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/New_York",
		*dbhost,
		*dbuser,
		*dbpwd,
		*dbname,
		*dbport,
	)

	db := OpenDatabaseConnection()

	if *shouldMigrateDB {
		log.Println("Skipping migrations")
		migrateDatabase(db)
	}
	return db
}

func OpenDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: configureLogger()})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&RemoteIndex{}, &MediaRequest{})
}

func configureLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}

func OpenRedisConnection() {

}
