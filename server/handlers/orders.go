package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/repositories"
)

func GetOrdersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := repositories.QueryAllOrders()
		if err != nil {
			fmt.Printf("Error querying database in GetItemsHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
