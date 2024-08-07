package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/repositories"
)

func GetItemsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := repositories.QueryAllItems()
		if err != nil {
			fmt.Printf("Error querying database in GetItemsHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func PostItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item = models.NewDefaultItem()
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			fmt.Printf("Error decoding JSON in PostItemHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		items, err := repositories.InsertItem(item)
		if err != nil {
			fmt.Printf("Error inserting item in PostItemHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error inserting item: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func PutItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var itemInput models.Item
		err := json.NewDecoder(r.Body).Decode(&itemInput)
		if err != nil {
			fmt.Printf("Error decoding JSON in PutItemHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		items, err := repositories.UpdateItem(&itemInput)
		if err != nil {
			fmt.Printf("Error updating item in PutItemHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error putting item: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}
