package db

import (
	"context"
	"fmt"
)

type Item struct {
	ID      int     `json:"id"`
	NameEng string  `json:"name_eng"`
	Price   float64 `json:"price"`
}

func GetItems() ([]Item, error) {
	rows, err := DBClient.Query(context.Background(), "SELECT id, name_eng, price FROM items")
	if err != nil {
		return nil, fmt.Errorf("failed to query items table: %w", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
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
