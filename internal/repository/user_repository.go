package repository

import (
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn // database dependency injection
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
