package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
)

func (r Repository) Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	const fn = "internal.storage.postgres.locks.Locks"
	query := `SELECT * FROM fn_locks_limit_offset($1, $2)`
	rows, err := r.pool.Query(ctx, query, pageNumber, recordsOnPage)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", fn, storage.ErrRecordsNotFound)
		}
	}
	defer rows.Close()
	recordsFromDB := make([]models.Lock, 0, recordsOnPage)
	for rows.Next() {
		scannedRow := models.Lock{}
		err = rows.Scan(
			&scannedRow.PartNumber,
			&scannedRow.Title,
			&scannedRow.Image,
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
