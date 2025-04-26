package repositories

import (
	"context"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
)

func QueryAllOrdersItems() ([]models.OrderItem, error) {
	query := "SELECT id, item_id, order_id, quantity, price FROM order_items"
	rows, err := db.DBClient.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all orders items: - Error: %v", err)
	}
	defer rows.Close()
	items, err := scanOrdersItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan orders items: %w", err)
	}
	return items, nil
}

func scanOrdersItems(rows pgx.Rows) ([]models.OrderItem, error) {
	defer rows.Close()
	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		err := rows.Scan(&item.ID, &item.ItemID, &item.OrderID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan orders items row: %w", err)
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return items, nil
}
