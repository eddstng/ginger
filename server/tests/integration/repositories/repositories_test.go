package repositories_test

import (
	"fmt"
	"os"
	"server/db"
	"server/models"
	"server/repositories"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
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

func TestQueryAllItems(t *testing.T) {
	items, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 18)
	require.Equal(t, 1, items[0].ID)
	require.Equal(t, 2, items[0].CategoryID)
	require.Equal(t, "Spring Rolls", items[0].NameEng)
	require.Equal(t, float64(5.99), items[0].Price)
}

func TestInsertItem(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = "TestInsertItem"
	testItem.Price = float64(99.99)

	items, err := repositories.InsertItem(testItem)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.IsType(t, 1, items[0].ID)
	require.Equal(t, "TestInsertItem", items[0].NameEng)
	require.Equal(t, float64(99.99), items[0].Price)
}

func TestInsertItemCategoryIdFKFail(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = "TestInsertItem"
	testItem.CategoryID = 9999

	items, err := repositories.InsertItem(testItem)
	require.Error(t, err)
	require.Equal(t, `failed to scan items: ERROR: insert or update on table "items" violates foreign key constraint "items_category_id_fkey" (SQLSTATE 23503)`, err.Error())
	require.Nil(t, items)
}

func TestInsertItemFailNameEngFail100Len(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = "TestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFail"

	items, err := repositories.InsertItem(testItem)
	require.Error(t, err)
	require.Equal(t, `failed to scan items: ERROR: value too long for type character varying(100) (SQLSTATE 22001)`, err.Error())
	require.Nil(t, items)
}

func TestUpdateItem(t *testing.T) {
	allItems, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.Len(t, allItems, 19)
	require.Equal(t, "TestInsertItem", allItems[18].NameEng)

	var testItemInput = models.NewDefaultItemInput()
	*testItemInput.ID = allItems[18].ID
	*testItemInput.NameEng = "TestUpdateItem"
	*testItemInput.Price = 11.15
	*testItemInput.NameEng = "TestUpdateItem"

	updatedItems, err := repositories.UpdateItem(testItemInput)
	require.NoError(t, err)
	require.NotNil(t, updatedItems)
	require.Len(t, updatedItems, 1)
	allItems, _ = repositories.QueryAllItems()
	require.Equal(t, "TestUpdateItem", allItems[18].NameEng)
	require.Equal(t, 11.15, allItems[18].Price)
}
