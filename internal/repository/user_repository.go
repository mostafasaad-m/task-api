package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/mostafasaad-m/task-api/internal/models"
)

type UserRepository struct {
	db *pgx.Conn //database dependency injection
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`

	return r.db.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}
