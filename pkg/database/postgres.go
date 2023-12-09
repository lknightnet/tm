package database

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(ctx context.Context, connURI string) (*Postgres, error) {
	pg := &Postgres{}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	cfg, err := pgxpool.ParseConfig(connURI)
	if err != nil {
		return nil, fmt.Errorf("postgres/new: parse connURI is failed: %s", err.Error())
	}

	cfg.MaxConns = 10

	pg.Pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres/new: connect with postgres is failed: %s", err.Error())
	}

	return pg, nil
}
