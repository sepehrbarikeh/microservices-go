package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt time.Time
}

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

func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User
	query := `SELECT id, email, password, created_at FROM users WHERE email=$1`
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return User{},err
	}
	fmt.Println(user)
	return user, nil
}
