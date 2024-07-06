package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var DBClient DBClientInterface

type DBClientInterface interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close(ctx context.Context) error
}

func SetDBClient(client DBClientInterface) {
	DBClient = client
}

func InitDBClientFromURL(databaseURL string) error {
	fmt.Println("Initializing database connection...")
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("failed to initialize DB client %s: %w", databaseURL, err)
	}
	SetDBClient(conn)
	fmt.Println("Database connected!")
	return nil
}

func CloseDBClient() error {
	if DBClient == nil {
		return nil
	}
	err := DBClient.Close(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close DB client: %w", err)
	}
	DBClient = nil
	fmt.Println("Disconnected from database!")
	return nil
}
