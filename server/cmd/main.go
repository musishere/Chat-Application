package main

import (
	"log"

	"github.com/musishere/chat-app/db"
)

func main() {
	_, err := db.NewDbConnection()
	if err != nil {
		log.Fatalf("Error initializing the database %s", err)
	}
}
