package integration_test

import (
	"os"
	"server/db"
	"server/repositories"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestQueryAllItems(t *testing.T) {
	err := godotenv.Load("../../.env")
	databaseURL := os.Getenv("DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
	defer db.CloseDBClient()
	items, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 9)
	// Bring these requirements back when we have a test database ready.
	// require.Equal(t, 1, items[0].ID)
	// require.Equal(t, "item1", items[0].NameEng)
	// require.Equal(t, float64(100), items[0].Price)
}
