package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgresql struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, url string) (*Postgresql, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Postgresql{
		db: pool,
	}, nil
}
