package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/repositories"
)

func GetItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repositories.QueryAllItems()
		if err != nil {
			fmt.Printf("Error in GetItemsHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error in GetItemsHandler: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}
