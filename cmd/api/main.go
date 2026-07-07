package main

import (
	"context"

	"log"

	"github.com/mostafasaad-m/task-api/internal/config"
	"github.com/mostafasaad-m/task-api/internal/models"
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
	repo := repository.NewUserRepository(db)

	user := &models.User{
		Username:     "mostafa",
		Email:        "mostafa@example.com",
		PasswordHash: "not_hashed_yet",
	}

	if err := repo.Create(user); err != nil {
		log.Fatal(err)
	}

	log.Printf("Created user with ID %d", user.ID)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	srv := server.New(cfg)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
