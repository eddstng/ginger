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

var SelectAllItemsQuery = "SELECT id, category_id, name_eng, price FROM items ORDER BY id ASC"

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

func InsertItem(item models.Item) ([]models.Item, error) {
	query := "INSERT INTO items (name_eng, price, category_id) VALUES ($1, $2, $3) RETURNING id, category_id, name_eng, price"
	rows, err := db.DBClient.Query(context.Background(), query, item.NameEng, item.Price, item.CategoryID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return nil, fmt.Errorf("failed to insert item: category ID %v does not exist (error code: %s)", item.CategoryID, pgErr.Code)
			default:
				return nil, fmt.Errorf("failed to insert item: %v (error code: %s)", pgErr.Message, pgErr.Code)
			}
		}
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}

func UpdateItem(item models.Item) ([]models.Item, error) {
	query := "UPDATE items SET name_eng = $1, price = $2, category_id = $3 WHERE id = $4 RETURNING id, category_id, name_eng, price"
	rows, err := db.DBClient.Query(context.Background(), query, item.NameEng, item.Price, item.CategoryID, item.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}
	defer rows.Close()
	items, err := scanItems(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}
	return items, nil
}
