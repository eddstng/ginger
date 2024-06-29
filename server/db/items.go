package db

import (
	"context"
	"fmt"
	"server/models"

	"github.com/jackc/pgx/v5"
)

type ItemsTable struct{}

func (ItemsTable) Query(contextBackground context.Context) (pgx.Rows, error) {
	rows, err := DBClient.Query(contextBackground, "SELECT id, name_eng, price FROM items")
	if err != nil {
		return nil, fmt.Errorf("failed to query items table: %w", err)
	}
	return rows, nil
}

func (ItemsTable) Scan(rows pgx.Rows) ([]models.Item, error) {
	defer rows.Close()
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.NameEng, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan items row: %w", err)
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return items, nil
}
