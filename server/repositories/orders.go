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

func UpdateOrder(orderInput *models.Order) ([]models.Order, error) {
	order, err := QueryOrder(*orderInput.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to query order: %v", err)
	}
	updateOrderFields(order, orderInput)

	query := "UPDATE orders SET subtotal = $1, total = $2, gst = $3, pst = $4, discount = $5, timestamp = $6, category = $10::order_category, void = $7, paid = $8, customizations = $9, customer_id = $11 WHERE id = $12 RETURNING id, subtotal, total, gst, pst, discount, timestamp, category, void, paid, customizations, customer_id"
	rows, err := db.DBClient.Query(context.Background(), query, *order.Subtotal, *order.Total, *order.GST, *order.PST, *order.Discount, *order.Timestamp, *order.Void, *order.Paid, *order.Customizations, *order.Category, *order.CustomerID, *order.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %v", err)
	}

	defer rows.Close()
	orders, err := scanOrders(rows)

	if err != nil {
		return nil, fmt.Errorf("failed to scan orders: %w", err)
	}

	return orders, nil
}

func updateOrderFields(order *models.Order, orderInput *models.Order) {
	if orderInput.ID != nil {
		order.ID = orderInput.ID
	}
	if orderInput.Subtotal != nil {
		order.Subtotal = orderInput.Subtotal
	}
	if orderInput.Total != nil {
		order.Total = orderInput.Total
	}
	if orderInput.GST != nil {
		order.GST = orderInput.GST
	}
	if orderInput.PST != nil {
		order.PST = orderInput.PST
	}
	if orderInput.Discount != nil {
		order.Discount = orderInput.Discount
	}
	if orderInput.Timestamp != nil {
		order.Timestamp = orderInput.Timestamp
	}
	if orderInput.Void != nil {
		order.Void = orderInput.Void
	}
	if orderInput.Paid != nil {
		order.Paid = orderInput.Paid
	}
	if orderInput.Customizations != nil {
		order.Customizations = orderInput.Customizations
	}
	if orderInput.Category != nil {
		order.Category = orderInput.Category
	}
	if orderInput.CustomerID != nil {
		order.CustomerID = orderInput.CustomerID
	}
}

func QueryOrder(id int) (*models.Order, error) {
	query := `SELECT id, subtotal, total, gst, pst, discount, timestamp, category, void, paid, customizations, customer_id 
    FROM orders 
    WHERE id = $1`
	rows, err := db.DBClient.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %v", err)
	}
	defer rows.Close()
	orders, err := scanOrders(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan orders: %w", err)
	}
	if len(orders) == 0 {
		return nil, nil
	}

	return &orders[0], nil
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
