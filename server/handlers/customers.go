package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/repositories"
)

func GetCustomersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := repositories.QueryAllCustomers()
		if err != nil {
			fmt.Printf("Error querying database in GetCustomersHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func PostCustomerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer = models.NewDefaultCustomer()
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			fmt.Printf("Error decoding JSON in PostCustomerHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		if *customer.Phone == "" && *customer.Name == "" {
			fmt.Printf("Name or Phone error: %v\n", err)
			http.Error(w, "Name or Phone are required fields", http.StatusBadRequest)
			return
		}
		customers, err := repositories.InsertCustomer(customer)
		if err != nil {
			fmt.Printf("Error inserting customer in PostCustomerHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error inserting customer: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func PutCustomerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customerInput models.Customer
		err := json.NewDecoder(r.Body).Decode(&customerInput)
		if err != nil {
			fmt.Printf("Error decoding JSON in PutCustomerHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		customers, err := repositories.UpdateCustomer(&customerInput)
		if err != nil {
			fmt.Printf("Error updating customer in PutCustomerHandler: %v\n", err)
			http.Error(w, fmt.Sprintf("Error putting customer: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
