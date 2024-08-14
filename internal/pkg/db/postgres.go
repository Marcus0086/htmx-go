package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres() (*pgxpool.Pool, error) {
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("POSTGRES_URL is not set")
	}

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
