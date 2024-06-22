package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/db"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := db.GetItems()
	if err != nil {
		var errorPrefix string = "Error in GetItemsHandler"
		http.Error(w, errorPrefix, http.StatusInternalServerError)
		log.Println(errorPrefix, err)
		return
	}

	jsonData, err := json.Marshal(items)
	if err != nil {
		var errorPrefix string = "Failed to marshal items to JSON"
		http.Error(w, errorPrefix, http.StatusInternalServerError)
		log.Println(errorPrefix, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
