package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mostafasaad-m/task-api/internal/config"
	"github.com/mostafasaad-m/task-api/internal/server"
	"github.com/mostafasaad-m/task-api/internal/store"
)

func main() {
	fmt.Println("Hello World")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	srv := server.New(cfg)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
