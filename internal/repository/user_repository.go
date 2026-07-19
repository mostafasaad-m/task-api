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

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	query := `
	SELECT
		id,
		username,
		email,
		password_hash,
		created_at,
		updated_at
	FROM users
	WHERE email = $1
	`

	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
