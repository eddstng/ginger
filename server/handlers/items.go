package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
)

var itemsTable = db.ItemsTable{}

func GetItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := itemsTable.Query(context.Background())
		if err != nil {
			fmt.Printf("Error in GetItemsHandler-Query: %v\n", err)
			http.Error(w, "Error in GetItemsHandler", http.StatusInternalServerError)
			return
		}
		items, err := itemsTable.Scan(rows)
		if err != nil {
			fmt.Printf("Error in GetItemsHandler-Scan: %v\n", err)
			http.Error(w, "Error in GetItemsHandler", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}
