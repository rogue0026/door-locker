package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
)

func (r Repository) GetByID(ctx context.Context, userID int64) (models.Account, error) {
	const fn = "internal.storage.postgres.accounts.GetByID"
	sqlQuery := `
SELECT 
	user_id,
	login,
	password_hash,
	status,
	first_name,
	last_name,
	birth_date,
	phone_mobile,
	email
FROM accounts
WHERE user_id = @user_id;
`
	args := pgx.NamedArgs{
		"user_id": userID,
	}
	account := models.Account{}
	err := r.pool.QueryRow(ctx, sqlQuery, args).Scan(&account.UserID,
		&account.Login,
		&account.Password,
		&account.Status,
		&account.FirstName,
		&account.LastName,
		&account.BirthDate,
		&account.PhoneMobile,
		&account.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return account, fmt.Errorf("%s: %w", fn, storage.ErrRecordsNotFound)
		}
	}
	return account, nil
}
