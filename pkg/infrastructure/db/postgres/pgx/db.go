package pgx

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConn(ctx context.Context, dataSourceName string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, dataSourceName)
}

func NewPool(ctx context.Context, dataSourceName string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, dataSourceName)
}
