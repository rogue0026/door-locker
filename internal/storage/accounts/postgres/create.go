package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) Save(ctx context.Context, account models.Account) error {
	const fn = "internal.storage.postgres.accounts.Save"
	err := account.EncryptPassword()
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	sqlQuery := `
call add_account(
	@login,
	@password_hash,
	@status,
	@first_name,
	@last_name,
	@birth_date,
	@phone_mobile,
	@email
);`
	args := pgx.NamedArgs{
		"login":         account.Login,
		"password_hash": account.Password,
		"status":        account.Status,
		"first_name":    account.FirstName,
		"last_name":     account.LastName,
		"birth_date":    account.BirthDate,
		"phone_mobile":  account.PhoneMobile,
		"email":         account.Email,
	}
	_, err = r.pool.Exec(ctx, sqlQuery, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
