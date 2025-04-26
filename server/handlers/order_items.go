package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/repositories"
)

func GetOrderItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetOrderItemsHandler")
		orders, err := repositories.QueryAllOrdersItems()
		if err != nil {
			fmt.Printf("Error querying database in GetOrdersHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
