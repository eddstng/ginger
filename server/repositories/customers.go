package repositories

import (
	"context"
	"errors"
	"fmt"
	"server/db"
	"server/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func QueryAllCustomers() ([]models.Customer, error) {
	var selectAllCustomersQuery = "SELECT id, name, phone, unit_number, street_number, street_name, buzzer_number, note FROM customers ORDER BY id ASC"
	rows, err := db.DBClient.Query(context.Background(), selectAllCustomersQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query all customers - Error: %v", err)
	}
	defer rows.Close()

	customers, err := scanCustomers(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan customers: %w", err)
	}
	return customers, nil
}

func InsertCustomer(customer *models.Customer) ([]models.Customer, error) {
	query := "INSERT INTO customers (name, phone, unit_number, street_number, street_name, buzzer_number, note) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, phone, unit_number, street_number, street_name, buzzer_number, note"
	rows, err := db.DBClient.Query(context.Background(), query, *customer.Name, *customer.Phone, *customer.UnitNumber, *customer.StreetNumber, *customer.StreetName, *customer.BuzzerNumber, *customer.Note)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				return nil, fmt.Errorf("failed to insert item: %v (error code: %s)", pgErr.Message, pgErr.Code)
			}
		}
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}
	defer rows.Close()
	customers, err := scanCustomers(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan customers: %w", err)
	}
	return customers, nil
}

func UpdateCustomer(customerInput *models.Customer) ([]models.Customer, error) {
	customer, err := QueryCustomer(*customerInput.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to query customer: %v", err)
	}
	updateCustomerFields(customer, customerInput)

	query := "UPDATE customers SET name = $1, phone = $2, unit_number = $3, street_number = $4, street_name = $5, buzzer_number = $6, note = $7 WHERE id = $8 RETURNING id, name, phone, unit_number, street_number, street_name, buzzer_number, note"
	rows, err := db.DBClient.Query(context.Background(), query, *customer.Name, *customer.Phone, *customer.UnitNumber, *customer.StreetNumber, *customer.StreetName, *customer.BuzzerNumber, *customer.Note, *customer.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to update item: %v", err)
	}

	defer rows.Close()
	customers, err := scanCustomers(rows)

	if err != nil {
		return nil, fmt.Errorf("failed to scan customers: %w", err)
	}

	return customers, nil
}

func QueryCustomer(id int) (*models.Customer, error) {
	query := "SELECT id, name, phone, unit_number, street_number, street_name, buzzer_number, note FROM customers WHERE id = $1"
	rows, err := db.DBClient.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query customer: %v", err)
	}
	defer rows.Close()
	customers, err := scanCustomers(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan customers: %w", err)
	}
	if len(customers) == 0 {
		return nil, nil
	}
	return &customers[0], nil
}

func updateCustomerFields(customer *models.Customer, customerInput *models.Customer) {
	if customerInput.ID != nil {
		customer.ID = customerInput.ID
	}
	if customerInput.Name != nil {
		customer.Name = customerInput.Name
	}
	if customerInput.Phone != nil {
		customer.Phone = customerInput.Phone
	}
	if customerInput.UnitNumber != nil {
		customer.UnitNumber = customerInput.UnitNumber
	}
	if customerInput.StreetNumber != nil {
		customer.StreetNumber = customerInput.StreetNumber
	}
	if customerInput.StreetName != nil {
		customer.StreetName = customerInput.StreetName
	}
	if customerInput.BuzzerNumber != nil {
		customer.BuzzerNumber = customerInput.BuzzerNumber
	}
	if customerInput.Note != nil {
		customer.Note = customerInput.Note
	}
}

func scanCustomers(rows pgx.Rows) ([]models.Customer, error) {
	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Phone,
			&customer.UnitNumber,
			&customer.StreetNumber,
			&customer.StreetName,
			&customer.BuzzerNumber,
			&customer.Note,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customers row: %w", err)
		}
		customers = append(customers, customer)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return customers, nil
}
