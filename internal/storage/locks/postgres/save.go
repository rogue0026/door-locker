package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) Save(ctx context.Context, lock models.Lock) error {
	const fn = "internal.storage.postgres.locks.Save"
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
	_, err := r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
