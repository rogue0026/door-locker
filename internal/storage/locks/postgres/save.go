package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) Save(ctx context.Context, lock models.Lock) error {
	const fn = "internal.storage.locks.postgres.Save"
	args := pgx.NamedArgs{
		"part_number":            lock.PartNumber,
		"title":                  lock.Title,
		"image":                  lock.Image,
		"price":                  lock.Price,
		"sale_price":             lock.SalePrice,
		"equipment":              lock.Equipment,
		"colors":                 lock.Colors,
		"description":            lock.Description,
		"category":               lock.Category,
		"card_memory":            lock.CardMemory,
		"material":               lock.Material,
		"has_mobile_application": lock.HasMobileApplication,
		"power_supply":           lock.PowerSupply,
		"size":                   lock.Size,
		"weight":                 lock.Weight,
		"door_type":              lock.DoorType,
		"door_thickness_min":     lock.DoorThicknessMin,
		"door_thickness_max":     lock.DoorThicknessMax,
		"rating":                 lock.Rating,
		"quantity":               lock.Quantity,
	}

	query := ``

	_, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
