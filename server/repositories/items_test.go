package repositories

import (
	"context"
	"server/db"
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
