package main

import (
	"log"

	"github.com/mostafasaad-m/task-api/internal/config"
	"github.com/mostafasaad-m/task-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(cfg)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
