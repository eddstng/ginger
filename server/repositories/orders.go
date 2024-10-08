package repositories

import (
	"context"
	"errors"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func InsertOrder(order *models.Order) ([]models.Order, error) {
	query := "INSERT INTO orders (subtotal, total, gst, pst, discount, timestamp, category, void, paid, customizations, customer_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, subtotal, total, gst, pst, discount, timestamp, category, void, paid, customizations, customer_id"
	rows, err := db.DBClient.Query(context.Background(), query, *order.Subtotal, *order.Total, *order.GST, *order.PST, *order.Discount, *order.Timestamp, *order.Category, *order.Void, *order.Paid, *order.Customizations, *order.CustomerID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				return nil, fmt.Errorf("failed to insert order: %v (error code: %s)", pgErr.Message, pgErr.Code)
			}
		}
		return nil, fmt.Errorf("failed to insert order: %v", err)
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
