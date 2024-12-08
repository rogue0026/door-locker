package postgres

import (
	"context"
	"fmt"

	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	const fn = "internal.storage.locks.postgres.Locks"
	query := `SELECT * FROM fn_locks_limit_offset($1, $2)`
	rows, err := r.pool.Query(ctx, query, pageNumber, recordsOnPage)
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
			&scannedRow.Images,
			&scannedRow.Price,
			&scannedRow.SalePrice,
			&scannedRow.Equipment,
			&scannedRow.Colors,
			&scannedRow.Description,
			&scannedRow.Category,
			&scannedRow.CardMemory,
			&scannedRow.Material,
			&scannedRow.HasMobileApplication,
			&scannedRow.PowerSupply,
			&scannedRow.Size,
			&scannedRow.Weight,
			&scannedRow.DoorType,
			&scannedRow.DoorThicknessMin,
			&scannedRow.DoorThicknessMax,
			&scannedRow.Rating,
			&scannedRow.Quantity)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		recordsFromDB = append(recordsFromDB, scannedRow)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return recordsFromDB, nil
}
