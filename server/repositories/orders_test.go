package repositories

import (
	"context"
	"server/db"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryAllOrdersWithMockedDB(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)

	test_helpers.MockGetOrdersQuery(mock)
	orders, err := QueryAllOrders()
	require.NoError(t, err)
	require.NotNil(t, orders)
	require.Len(t, orders, 2)
	require.Equal(t, 1, *orders[0].ID)
	expectedOrders := []models.Order{
		{
			ID:             models.PtrInt(1),
			Subtotal:       models.PtrFloat64(7.50),
			Total:          models.PtrFloat64(7.87),
			GST:            models.PtrFloat64(0.37),
			PST:            models.PtrFloat64(0.00),
			Discount:       models.PtrFloat64(0.00),
			Category:       models.PtrString("IN"),
			Customizations: nil,
			CustomerID:     models.PtrInt(1),
		},
		{
			ID:             models.PtrInt(2),
			Subtotal:       models.PtrFloat64(6.00),
			Total:          models.PtrFloat64(6.30),
			GST:            models.PtrFloat64(0.30),
			PST:            models.PtrFloat64(0.00),
			Discount:       models.PtrFloat64(0.00),
			Category:       models.PtrString("OUT"),
			Customizations: models.PtrString(`[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]`),
			CustomerID:     models.PtrInt(2),
		},
	}
	require.Equal(t, expectedOrders[0], orders[0])
	require.Equal(t, expectedOrders[1], orders[1])
}
