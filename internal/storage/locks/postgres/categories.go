package postgres

import (
	"context"
	"fmt"
	"github.com/rogue0026/door-locker/internal/models"
)

func (r Repository) Categories(ctx context.Context) ([]models.Category, error) {
	const fn = "internal.storage.postgres.locks.Categories"

	query := `
SELECT
	id,
	name, 
	image
FROM lock_categories;`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()
	records := make([]models.Category, 0)
	for rows.Next() {
		currentRecord := models.Category{}
		err = rows.Scan(&currentRecord.ID, &currentRecord.Name, &currentRecord.Image)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		records = append(records, currentRecord)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return records, nil
}
