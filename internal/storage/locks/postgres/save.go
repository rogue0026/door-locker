package postgres

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rogue0026/door-locker/internal/models"
)

var (
	ErrDecodingImage = errors.New("error while decoding image")
)

func (r Repository) Save(ctx context.Context, lock models.Lock) error {
	const fn = "internal.storage.locks.postgres.Save"
	filenames, err := decodeImages(lock.Title, lock.Images)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	lock.Images = make([]string, len(filenames))
	copy(lock.Images, filenames)
	args := pgx.NamedArgs{
		"title":                  lock.Title,
		"image":                  lock.Images,
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

	query := `
call save_lock(
@title,
@image,
@price,
@sale_price,
@equipment,
@colors,
@description,
@category,
@card_memory,
@material,
@has_mobile_application,
@power_supply,
@size,
@weight,
@door_type,
@door_thickness_min,
@door_thickness_max,
@rating,
@quantity);`

	_, err = r.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}

func decodeImages(title string, images []string) ([]string, error) {
	_, err := os.Stat("./images")
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir("./images", 0777)
		if err != nil {
			return nil, err
		}
	}
	fileNames := make([]string, 0)
	for i := 0; i < len(images); i++ {
		elems := strings.Split(images[i], ",")
		if len(elems) != 2 {
			return nil, fmt.Errorf("%w: bad image data", ErrDecodingImage)
		}
		if strings.Contains(elems[0], "png") {
			decodedImage, err := base64.StdEncoding.DecodeString(elems[1])
			if err != nil {
				return nil, fmt.Errorf("%w: %w", ErrDecodingImage, err)
			}
			title = strings.Replace(title, " ", "_", -1)
			var filename = fmt.Sprintf("%d_%s.png", time.Now().UnixMilli(), title)
			f, err := os.Create(fmt.Sprintf("./images/%s", filename))
			if err != nil {
				return nil, err
			}
			if _, err = f.Write(decodedImage); err != nil {
				return nil, err
			}
			if err = f.Sync(); err != nil {
				return nil, err
			}
			if err = f.Close(); err != nil {
				return nil, err
			}
			fileNames = append(fileNames, filename)
		}
	}
	return fileNames, nil
}
