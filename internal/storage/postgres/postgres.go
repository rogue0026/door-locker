package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogue0026/door-locker/internal/models"
)

type Storage struct {
	connPool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Storage, error) {
	const fn = "internal.storage.postgres.New"
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	s := Storage{
		connPool: pool,
	}
	return &s, nil
}

func (s Storage) Locks(ctx context.Context, limit uint, offset uint) ([]models.DoorLock, error) {
	const fn = "internal.storage.postgres.Locks"

}
