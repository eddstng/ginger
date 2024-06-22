package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var DBClient *pgx.Conn

func InitDBClient(databaseURL string) error {
	fmt.Println("Initializing database connection...")
	var err error
	DBClient, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("failed to initialize DB client%s: %w", databaseURL, err)
	}
	fmt.Println("Database connected!")
	return nil
}

func CloseDBClient() error {
	err := DBClient.Close(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close DB client: %w", err)
	}
	DBClient = nil
	fmt.Println("Disconnected from database!")
	return nil
}
