package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbURL string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("db ping failed:", err)
	}

	fmt.Println("✅ connected to postgres")

		err = Migrate(pool)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
