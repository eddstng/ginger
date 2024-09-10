package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/repositories"
)

func GetOrdersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := repositories.QueryAllOrders()
		if err != nil {
			fmt.Printf("Error querying database in GetOrdersHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}

func PostOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order = models.NewDefaultOrder()
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Printf("Error decoding JSON in PostOrderHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		orders, err := repositories.InsertOrder(order)
		if err != nil {
			fmt.Printf("Error inserting order in PostOrderHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error inserting order: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
