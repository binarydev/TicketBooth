package datastore

import (
	"fmt"
	"github.com/binarydev/ticketbooth/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresDsn string
var SkipMigrateDB *bool

func InitializeDatabase() {

	db := OpenDatabaseConnection()

	if *SkipMigrateDB {
		log.Println("Skipping migrations")
	} else {
		log.Println("Migrating databases")
		migrateDatabase(db)
	}
	fmt.Println("Database connection successfully tested")
}

func OpenDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(PostgresDsn), &gorm.Config{Logger: configureLogger()})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.RemoteIndex{}, &models.MediaRequest{})
	if err != nil {
		log.Fatal(err)
		return
	}
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
