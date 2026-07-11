package main

import (
	"context"

	"log"

	"github.com/mostafasaad-m/task-api/internal/config"
	"github.com/mostafasaad-m/task-api/internal/handlers"
	"github.com/mostafasaad-m/task-api/internal/repository"

	"github.com/mostafasaad-m/task-api/internal/server"
	"github.com/mostafasaad-m/task-api/internal/store"
)

func main() {
	//fmt.Println("ci/cd last check")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	repo := repository.NewUserRepository(db)

	authHandler := handlers.NewAuthHandler(repo)

	srv := server.New(cfg, authHandler)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
