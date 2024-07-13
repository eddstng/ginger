package repositories

import (
	"context"
	"server/db"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryAllItemsWithMockedDB(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockGetItemsQuery(mock)
	items, err := QueryAllItems()
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, 1, items[0].ID)
	require.Equal(t, "item1", items[0].NameEng)
	require.Equal(t, float64(100), items[0].Price)
}

func TestInsertItemWithMockedDB(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertItemQuery(mock, models.Item{ID: 111, NameEng: "test_item_insert", Price: 111, CategoryID: 111})
	items, err := InsertItem(models.Item{NameEng: "test_item_insert", Price: 111, CategoryID: 111})
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, models.Item{ID: 111, NameEng: "test_item_insert", Price: 111, CategoryID: 111}, items[0])
}

func TestUpdateItemWithMockedDB(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockUpdateItemQuery(mock, models.Item{ID: 222, NameEng: "test_item_update", Price: 222, CategoryID: 222})
	items, err := UpdateItem(models.Item{ID: 222, NameEng: "test_item_update", Price: 222, CategoryID: 222})
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, models.Item{ID: 222, NameEng: "test_item_update", Price: 222, CategoryID: 222}, items[0])
}
