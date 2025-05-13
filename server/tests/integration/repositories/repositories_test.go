package repositories_test

import (
	"fmt"
	"os"
	"server/db"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Failed to load environment variables:", err)
		os.Exit(1)
	}

	databaseURL := os.Getenv("TEST_DATABASE_URL")
	err = db.InitDBClientFromURL(databaseURL)
	if err != nil {
		fmt.Println("Failed to initialize DB client:", err)
		os.Exit(1)
	}

	if db.DBClient == nil {
		fmt.Println("DB client is nil")
		os.Exit(1)
	}

	code := m.Run()
	db.CloseDBClient()
	os.Exit(code)
}
