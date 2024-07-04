package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
)

func GetItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := GetItems("SELECT id, name_eng, price FROM items")
		if err != nil {
			fmt.Printf("Error in GetItemsHandler: %v\n", err)
			http.Error(w, "Error in GetItemsHandler", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func GetItems(query string) ([]models.Item, error) {
	rows, err := db.DBClient.Query(context.Background(), query)
	if err != nil {
		fmt.Printf("Error in GetItems: %v\n", err)
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
