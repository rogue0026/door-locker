package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
)

type Storage struct {
	connPool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (Storage, error) {
	const fn = "internal.storage.postgres.New"
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return Storage{}, fmt.Errorf("%s: %w", fn, err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		return Storage{}, fmt.Errorf("%s: %w", fn, err)
	}
	s := Storage{
		connPool: pool,
	}
	return s, nil
}

func (s Storage) LocksWithLimitOffset(ctx context.Context, recordsOnPage int64, pageNumber int64) ([]models.DoorLock, error) {
	const fn = "internal.storage.postgres.Locks"
	query := `SELECT * FROM door_locks ORDER BY part_number OFFSET $1 LIMIT $2`
	var offset = (pageNumber - 1) * recordsOnPage
	rows, err := s.connPool.Query(ctx, query, offset, recordsOnPage)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", fn, storage.ErrRecordsNotFound)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()
	recordsFromDB := make([]models.DoorLock, 0, recordsOnPage)
	for rows.Next() {
		scannedRow := models.DoorLock{}
		err = rows.Scan(
			&scannedRow.PartNumber,
			&scannedRow.Title,
			&scannedRow.Price,
			&scannedRow.SalePrice,
			&scannedRow.Equipment,
			&scannedRow.ColorID,
			&scannedRow.Description,
			&scannedRow.CategoryID,
			&scannedRow.CardMemory,
			&scannedRow.MaterialID,
			&scannedRow.HasMobileApplication,
			&scannedRow.PowerSupply,
			&scannedRow.Size,
			&scannedRow.Weight,
			&scannedRow.DoorsTypeID,
			&scannedRow.DoorThicknessMin,
			&scannedRow.DoorThicknessMax,
			&scannedRow.Rating,
			&scannedRow.Quantity)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		recordsFromDB = append(recordsFromDB, scannedRow)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return recordsFromDB, nil
}

func (s Storage) Ping(ctx context.Context) error {
	return s.connPool.Ping(ctx)
}
