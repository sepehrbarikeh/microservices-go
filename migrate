package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(db *pgxpool.Pool) error {

	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		product TEXT NOT NULL,
		quantity INT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`
	_, err := db.Exec(context.Background(), query)
	return err
}
