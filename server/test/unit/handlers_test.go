package unit_test

import (
	"context"
	"server/db"
	"server/handlers"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
)

func TestGetItems(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	mock.ExpectQuery("SELECT id, name_eng, price FROM items").WillReturnRows(mock.NewRows([]string{"id", "name_eng", "price"}).AddRow(1, "item1", float64(100)))
	items, err := handlers.GetItems("SELECT id, name_eng, price FROM items")
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, 1, items[0].ID)
	require.Equal(t, "item1", items[0].NameEng)
	require.Equal(t, float64(100), items[0].Price)
}
