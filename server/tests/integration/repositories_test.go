package integration_test

import (
	"os"
	"server/db"
	"server/models"
	"server/repositories"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestQueryAllItems(t *testing.T) {
	err := godotenv.Load("../../.env")
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
	defer db.CloseDBClient()
	items, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 6)
	// Bring these requirements back when we have a test database ready.
	// require.Equal(t, 1, items[0].ID)
	// require.Equal(t, "item1", items[0].NameEng)
	// require.Equal(t, float64(100), items[0].Price)
}

func TestInsertItem(t *testing.T) {
	err := godotenv.Load("../../.env")
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
	defer db.CloseDBClient()
	items, err := repositories.InsertItem(models.Item{NameEng: "test_items_insert", Price: 111, CategoryID: 1})
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	// require.Equal(t, 111, items[0].ID)
	require.IsType(t, 1, items[0].ID)
	require.Equal(t, "test_items_insert", items[0].NameEng)
	require.Equal(t, float64(111), items[0].Price)
}

func TestUpdateItem(t *testing.T) {
	err := godotenv.Load("../../.env")
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
	defer db.CloseDBClient()
	allItems, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.Equal(t, models.Item{ID: 2, NameEng: "Soup 1", Price: 3, CategoryID: 2}, allItems[1])
	updatedItems, err := repositories.UpdateItem(models.Item{ID: 2, NameEng: "test_item_update", Price: 222, CategoryID: 3})
	require.NoError(t, err)
	require.NotNil(t, updatedItems)
	require.Len(t, updatedItems, 1)
	require.Equal(t, models.Item{ID: 2, NameEng: "test_item_update", Price: 222, CategoryID: 3}, updatedItems[0])
	allItems, _ = repositories.QueryAllItems()
	require.Equal(t, models.Item{ID: 2, NameEng: "test_item_update", Price: 222, CategoryID: 3}, allItems[1])
}
