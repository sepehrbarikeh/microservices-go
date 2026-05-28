package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(email string, password string) (int, error) {
	var id int

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.db.QueryRow(context.Background(), query, email, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}