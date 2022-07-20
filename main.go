package main

import (
	"log"

	"github.com/binarydev/ticketbooth/datastore"
	"github.com/binarydev/ticketbooth/router"
)

func main() {
	log.SetFlags(log.Ldate | log.LUTC)
	datastore.InitializeDatabase()
	router.Engine.Run(":8080")
}
