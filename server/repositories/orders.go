package repositories

import (
	"context"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
)

func QueryAllOrders() ([]models.Order, error) {
	query := "SELECT id, subtotal, total, gst, pst, discount, timestamp, category, void, paid, customizations, customer_id FROM orders"
	rows, err := db.DBClient.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all orders - Error: %v", err)
	}
	defer rows.Close()

	orders, err := scanOrders(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan orders: %w", err)
	}
	return orders, nil
}

func scanOrders(rows pgx.Rows) ([]models.Order, error) {
	defer rows.Close()
	var orders []models.Order

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Subtotal, &order.Total, &order.GST, &order.PST, &order.Discount, &order.Timestamp, &order.Category, &order.Void, &order.Paid, &order.Customizations, &order.CustomerID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan orders row: %w", err)
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return orders, nil
}
