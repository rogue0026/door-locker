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

func (s Storage) LocksWithLimitOffset(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.DoorLock, error) {
	const fn = "internal.storage.postgres.LocksWithLimitOffset"
	query := `SELECT * from fn_locks_limit_offset($1, $2)`
	rows, err := s.connPool.Query(ctx, query, pageNumber, recordsOnPage)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", fn, storage.ErrRecordsNotFound)
		}
	}
	defer rows.Close()
	recordsFromDB := make([]models.DoorLock, 0, recordsOnPage)
	for rows.Next() {
		scannedRow := models.DoorLock{}
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

func (s Storage) Ping(ctx context.Context) error {
	return s.connPool.Ping(ctx)
}

func (s Storage) Close() {
	s.connPool.Close()
}

func (s Storage) SaveLock(ctx context.Context, lock models.DoorLock) error {
	const fn = "internal.storage.postgres.SaveLock"
	query := `
call save_door_lock(
@part_number,
@title,
@image,
@price,
@sale_price,
@equipment,
@color_id,
@description,
@category_id,
@card_memory,
@material_id,
@has_mobile_application,
@power_supply,
@size,
@weight,
@door_types_id,
@door_thickness_min,
@door_thickness_max,
@rating,
@quantity);`

	args := pgx.NamedArgs{
		"part_number":            lock.PartNumber,
		"title":                  lock.Title,
		"image":                  lock.Image,
		"price":                  lock.Price,
		"sale_price":             lock.SalePrice,
		"equipment":              lock.Equipment,
		"color_id":               lock.ColorID,
		"description":            lock.Description,
		"category_id":            lock.CategoryID,
		"card_memory":            lock.CardMemory,
		"material_id":            lock.MaterialID,
		"has_mobile_application": lock.HasMobileApplication,
		"power_supply":           lock.PowerSupply,
		"size":                   lock.Size,
		"weight":                 lock.Weight,
		"door_types_id":          lock.DoorsTypeID,
		"door_thickness_min":     lock.DoorThicknessMin,
		"door_thickness_max":     lock.DoorThicknessMax,
		"rating":                 lock.Rating,
		"quantity":               lock.Quantity,
	}
	_, err := s.connPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}

func (s Storage) DeleteLockByPartNumber(ctx context.Context, partNumber string) error {
	const fn = "internal.storage.postgres.DeleteLockByPartNumber"
	query := `call delete_door_lock_by_partnumber(@partnumber)`
	args := pgx.NamedArgs{
		"partnumber": partNumber,
	}
	_, err := s.connPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
