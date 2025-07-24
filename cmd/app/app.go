package app

import (
	"log"

	"github.com/Brian-M-J/social-media-app-go/internals/database"
	"github.com/Brian-M-J/social-media-app-go/internals/server"
)

func Setup() {
	database.Config()

	server.Setup()
	app := server.New()

	if app == nil {
		log.Fatalln("app is nil")
	}

	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
