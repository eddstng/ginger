package repositories

import (
	"context"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
)

var SelectAllItemsQuery = "SELECT id, category_id, name_eng, price FROM items"

func QueryAllItems() ([]models.Item, error) {
	rows, err := db.DBClient.Query(context.Background(), SelectAllItemsQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query all items - Error: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}

func scanItems(rows pgx.Rows) ([]models.Item, error) {
	defer rows.Close()
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.CategoryID, &item.NameEng, &item.Price)
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
