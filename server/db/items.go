package db

import (
	"context"
	"fmt"
	"server/models"
)

func GetItems() ([]models.Item, error) {
	rows, err := DBClient.Query(context.Background(), "SELECT id, name_eng, price FROM items")
	if err != nil {
		return nil, fmt.Errorf("failed to query items table: %w", err)
	}
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
