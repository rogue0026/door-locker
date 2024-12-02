package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (r Repository) Remove(ctx context.Context, userID int64) error {
	const fn = "internal.storage.postgres.accounts.Remove"
	args := pgx.NamedArgs{
		"user_id": userID,
	}
	query := `DELETE FROM accounts WHERE user_id = @user_id`
	_, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
