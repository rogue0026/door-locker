package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (r Repository) Delete(ctx context.Context, partNumber string) error {
	const fn = "internal.storage.postgres.locks.Remove"
	query := `call delete_door_lock_by_partnumber(@partnumber)`
	args := pgx.NamedArgs{
		"partnumber": partNumber,
	}
	_, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
