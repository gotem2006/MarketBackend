package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnPostgres() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:marta2010@localhost:5432/databaselab3")
	return conn, err
}
