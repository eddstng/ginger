package db_test

import (
	"os"
	"server/db"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestInitDBClientFromURLAndCloseDBClient(t *testing.T) {
	err := godotenv.Load("../../../.env")
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
	err = db.CloseDBClient()
	require.NoError(t, err)
}

func TestInitDBClientFromURLErrors(t *testing.T) {
	t.Run("URL Parse Error", func(t *testing.T) {
		os.Setenv("DATABASE_URL_MOCK", "mock_database_url")
		databaseURL := os.Getenv("DATABASE_URL_MOCK")
		err := db.InitDBClientFromURL(databaseURL)
		require.Contains(t, err.Error(), "failed to initialize DB client mock_database_url: cannot parse `mock_database_url`: failed to parse as keyword/value (invalid keyword/value)")
	})
	t.Run("DB Not Running at URL", func(t *testing.T) {
		os.Setenv("DATABASE_URL_MOCK", "postgres://postgres:postgres@localhost:3000/postgres?sslmode=disable")
		databaseURL := os.Getenv("DATABASE_URL_MOCK")
		err := db.InitDBClientFromURL(databaseURL)
		require.Contains(t, err.Error(), "dial error: dial tcp 127.0.0.1:3000: connect: connection refused")
	})
}
