package handlers

import (
	"encoding/json"
	"net/http"
	"server/db"
)

func GetItemsHandler(dbGetter db.DBGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := dbGetter.GetItems()
		if err != nil {
			http.Error(w, "Error in GetItemsHandler", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}
