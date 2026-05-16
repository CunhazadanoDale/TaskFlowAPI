package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(conn string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}