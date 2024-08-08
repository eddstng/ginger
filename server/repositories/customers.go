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
	fmt.Println(customer)
	query := "INSERT INTO customers (name, phone, unit_number, street_number, street_name, buzzer_number, note) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, phone, unit_number, street_number, street_name, buzzer_number, note"
	rows, err := db.DBClient.Query(context.Background(), query, *customer.Name, *customer.Phone, *customer.UnitNumber, *customer.StreetNumber, *customer.StreetName, *customer.BuzzerNumber, *customer.Note)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return nil, fmt.Errorf("failed to insert item: category ID %v does not exist (error code: %s)", customer.Phone, pgErr.Code)
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
