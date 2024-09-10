package repositories

import (
	"context"
	"server/db"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"
	"time"

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

func TestInsertOrderWithMockedDB(t *testing.T) {
	var testOrder = models.NewDefaultOrder()
	testOrder.Subtotal = models.PtrFloat64(7.50)
	testOrder.Total = models.PtrFloat64(7.87)
	testOrder.GST = models.PtrFloat64(0.37)
	testOrder.PST = models.PtrFloat64(0.00)
	testOrder.Discount = models.PtrFloat64(0.00)
	testOrder.Category = models.PtrString("IN")
	testOrder.CustomerID = models.PtrInt(1)
	testOrder.ID = models.PtrInt(30)
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertOrderQuery(mock, *testOrder)

	orders, err := InsertOrder(testOrder)
	require.NoError(t, err)
	require.NotNil(t, orders)
	require.Len(t, orders, 1)
	expectedOrders := []models.Order{
		{
			ID:             models.PtrInt(30),
			Subtotal:       models.PtrFloat64(7.50),
			Total:          models.PtrFloat64(7.87),
			GST:            models.PtrFloat64(0.37),
			PST:            models.PtrFloat64(0.00),
			Discount:       models.PtrFloat64(0.00),
			Category:       models.PtrString("IN"),
			Customizations: models.PtrString("[]"),
			CustomerID:     models.PtrInt(1),
			Timestamp:      new(time.Time),
			Void:           models.PtrBool(false),
			Paid:           models.PtrBool(false),
		},
	}
	require.Equal(t, expectedOrders[0], orders[0])
}
