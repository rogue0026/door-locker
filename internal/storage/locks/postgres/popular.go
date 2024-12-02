package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error) {
	const fn = "internal.storage.postgres.locks.LocksByRating"
	args := pgx.NamedArgs{
		"records_on_page": recordsOnPage,
	}
	query := `SELECT * FROM fn_locks_ordered_by_rating(@records_on_page);`
	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
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
		err = rows.Err()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
	}
	return recordsFromDB, nil
}
