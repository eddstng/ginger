package repositories_test

import (
	"encoding/json"
	"reflect"
	"server/models"
	"server/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQueryAllOrders(t *testing.T) {
	orders, err := repositories.QueryAllOrders()
	require.NoError(t, err)
	require.NotNil(t, orders)
	require.Len(t, orders, 2)

	require.Equal(t, 7.50, *orders[0].Subtotal)
	require.Equal(t, 7.87, *orders[0].Total)
	require.Equal(t, 0.37, *orders[0].GST)
	require.Equal(t, 0.00, *orders[0].PST)
	require.Equal(t, 0.00, *orders[0].Discount)
	require.Equal(t, "IN", *orders[0].Category)
	require.Nil(t, orders[0].Customizations)
	require.Equal(t, 1, *orders[0].CustomerID)

	require.Equal(t, 6.00, *orders[1].Subtotal)
	require.Equal(t, 6.30, *orders[1].Total)
	require.Equal(t, 0.30, *orders[1].GST)
	require.Equal(t, 0.00, *orders[1].PST)
	require.Equal(t, 0.00, *orders[1].Discount)
	require.Equal(t, "OUT", *orders[1].Category)
	require.NotNil(t, orders[1].Customizations)

	var expectedCustomizations, actualCustomizations []map[string]interface{}
	json.Unmarshal([]byte(`[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]`), &expectedCustomizations)
	json.Unmarshal([]byte(*orders[1].Customizations), &actualCustomizations)
	require.Equal(t, expectedCustomizations, actualCustomizations)

	require.Equal(t, 2, *orders[1].CustomerID)
}

func TestInsertOrder(t *testing.T) {
	var testOrder = models.NewDefaultOrder()
	testOrder.CustomerID = models.PtrInt(1)
	*testOrder.Subtotal = 7.50
	*testOrder.Total = 7.87
	*testOrder.GST = 0.37
	*testOrder.PST = 0.00
	*testOrder.Discount = 0.00
	*testOrder.Timestamp = time.Now()
	*testOrder.Void = false
	*testOrder.Paid = true
	*testOrder.Customizations = ""
	*testOrder.Customizations = `[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]`
	*testOrder.Category = "IN"
	testOrder.CustomerID = models.PtrInt(1)

	// the handler handles the nil values and turns it them into default values so we just use the NewDefaultOrder here is ok.
	orders, err := repositories.InsertOrder(testOrder)
	require.NoError(t, err)
	require.NotNil(t, orders)
	require.Len(t, orders, 1)
	require.IsType(t, *testOrder.ID, *orders[0].ID)
	require.Equal(t, *testOrder.Subtotal, *orders[0].Subtotal)
	require.Equal(t, *testOrder.Total, *orders[0].Total)
	require.Equal(t, *testOrder.GST, *orders[0].GST)
	require.Equal(t, *testOrder.PST, *orders[0].PST)
	require.Equal(t, *testOrder.Discount, *orders[0].Discount)
	require.Equal(t, "*time.Time", reflect.TypeOf(orders[0].Timestamp).String())
	require.Equal(t, *testOrder.Void, *orders[0].Void)
	require.Equal(t, *testOrder.Paid, *orders[0].Paid)

	var expectedCustomizations map[string]interface{}
	var actualCustomizations map[string]interface{}
	json.Unmarshal([]byte(*testOrder.Customizations), &expectedCustomizations)
	json.Unmarshal([]byte(*orders[0].Customizations), &actualCustomizations)

	require.Equal(t, *testOrder.Category, *orders[0].Category)
	require.Equal(t, *testOrder.CustomerID, *orders[0].CustomerID)
}

func TestUpdateOrder(t *testing.T) {
	allOrders, err := repositories.QueryAllOrders()
	require.NoError(t, err)
	require.Len(t, allOrders, 3)

	testOrder := allOrders[2]
	require.Equal(t, 7.50, *testOrder.Subtotal)
	require.Equal(t, 7.87, *testOrder.Total)
	require.Equal(t, 0.37, *testOrder.GST)
	require.Equal(t, 0.00, *testOrder.PST)
	require.Equal(t, 0.00, *testOrder.Discount)
	// require.Equal(t, time.Now(), *testOrder.Timestamp)
	require.WithinDuration(t, time.Now(), *testOrder.Timestamp, time.Second)
	require.Equal(t, false, *testOrder.Void)
	require.Equal(t, true, *testOrder.Paid)
	require.JSONEq(t, `[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]`, *testOrder.Customizations)
	require.Equal(t, "IN", *testOrder.Category)
	require.Equal(t, 1, *testOrder.CustomerID)

	var testOrderInput = models.NewDefaultOrderWithNil()
	testOrderInput.ID = allOrders[2].ID
	testOrderInput.Category = models.PtrString("OUT")
	testOrderInput.Customizations = models.PtrString(`[{"name_eng": "add ss sauce", "name_oth": "gulojup", "price": 1.00}]`)

	updatedOrder, err := repositories.UpdateOrder(testOrderInput)
	require.NoError(t, err)
	require.NotNil(t, updatedOrder)
	require.Len(t, updatedOrder, 1)
	allOrders, _ = repositories.QueryAllOrders()
	require.JSONEq(t, `[{"name_eng": "add ss sauce", "name_oth": "gulojup", "price": 1.00}]`, *allOrders[2].Customizations)
	require.Equal(t, "OUT", *allOrders[2].Category)
}
