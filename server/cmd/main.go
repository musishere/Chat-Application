package main

import (
	"log"

	"github.com/musishere/chat-app/db"
	"github.com/musishere/chat-app/internal/router"
	"github.com/musishere/chat-app/internal/user"
	"github.com/musishere/chat-app/internal/ws"
)

func main() {
	dbConn, err := db.NewDbConnection()
	if err != nil {
		log.Fatalf("Error initializing the database %s", err)
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	// http://localhost:8000/signup
	router.InitRouter(*userHandler, wsHandler)
	router.Start("0.0.0.0:8000")
}
