package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DBClient *pgx.Conn

func InitDBClient() error {
	fmt.Println("Initializing database connection...")
	var err error
	databaseURL := os.Getenv("DATABASE_URL")
	DBClient, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("failed to initialize DB client%s: %w", databaseURL, err)
	}
	fmt.Println("Database connected!")
	return nil
}

func CloseDBClient() {
	DBClient.Close(context.Background())
	fmt.Println("Disconnected from database!")
}
