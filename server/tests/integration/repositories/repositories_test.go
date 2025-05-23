package repositories_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"server/db"
	"server/models"
	"server/repositories"
	"testing"
	"time"

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
	require.Equal(t, 1, *items[0].ID)
	require.Equal(t, 2, *items[0].CategoryID)
	require.Equal(t, "Spring Rolls", *items[0].NameEng)
	require.Equal(t, float64(5.99), *items[0].Price)
}

func TestInsertItem(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = models.PtrString("TestInsertItem")
	testItem.Price = models.PtrFloat64(99.99)

	// the handler handles the item with nil values and turns it them into default values so we just use the NewDefaultItem here is ok.
	items, err := repositories.InsertItem(testItem)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.IsType(t, 1, *items[0].ID)
	require.Equal(t, "TestInsertItem", *items[0].NameEng)
	require.Equal(t, float64(99.99), *items[0].Price)
}

func TestInsertItemCategoryIdFKFail(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = models.PtrString("TestInsertItem")
	testItem.CategoryID = models.PtrInt(9999)

	items, err := repositories.InsertItem(testItem)
	require.Error(t, err)
	require.Equal(t, `failed to scan items: ERROR: insert or update on table "items" violates foreign key constraint "items_category_id_fkey" (SQLSTATE 23503)`, err.Error())
	require.Nil(t, items)
}

func TestInsertItemFailNameEngFail100Len(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = models.PtrString("TestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFailTestInsertItemFail")

	items, err := repositories.InsertItem(testItem)
	require.Error(t, err)
	require.Equal(t, `failed to scan items: ERROR: value too long for type character varying(100) (SQLSTATE 22001)`, err.Error())
	require.Nil(t, items)
}

func TestUpdateItem(t *testing.T) {
	allItems, err := repositories.QueryAllItems()
	require.NoError(t, err)
	require.Len(t, allItems, 19)
	require.Equal(t, "TestInsertItem", *allItems[18].NameEng)

	var testItemInput = models.NewDefaultItemWithNil()
	testItemInput.ID = allItems[18].ID
	testItemInput.NameEng = models.PtrString("TestUpdateItem")
	testItemInput.Price = models.PtrFloat64(11.15)
	testItemInput.NameEng = models.PtrString("TestUpdateItem")

	updatedItems, err := repositories.UpdateItem(testItemInput)
	require.NoError(t, err)
	require.NotNil(t, updatedItems)
	require.Len(t, updatedItems, 1)
	allItems, _ = repositories.QueryAllItems()
	require.Equal(t, "TestUpdateItem", *allItems[18].NameEng)
	require.Equal(t, 11.15, *allItems[18].Price)
}

func TestQueryAllCustomers(t *testing.T) {
	customers, err := repositories.QueryAllCustomers()
	require.NoError(t, err)
	require.NotNil(t, customers)
	require.Len(t, customers, 3)
	require.Equal(t, 1, *customers[0].ID)
	require.Equal(t, "John Doe", *customers[0].Name)
	require.Equal(t, "604-123-1234", *customers[0].Phone)
	require.Equal(t, "5555", *customers[0].StreetNumber)
	require.Equal(t, "Powel St", *customers[0].StreetName)
}

func TestInsertCustomer(t *testing.T) {
	var testCustomer = models.NewDefaultCustomer()
	// testCustomer.Name = models.PtrString("TestInsertCustomer")
	// testCustomer.Phone = models.PtrString("604-333-3838")
	*testCustomer.Name = "TestInsertCustomer"
	*testCustomer.Phone = "604-333-3838"
	items, err := repositories.InsertCustomer(testCustomer)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.IsType(t, 1, *items[0].ID)
	require.Equal(t, "TestInsertCustomer", *items[0].Name)
	require.Equal(t, "604-333-3838", *items[0].Phone)
	require.Equal(t, "", *items[0].StreetName)
	require.Equal(t, "", *items[0].StreetNumber)
	require.Equal(t, "", *items[0].BuzzerNumber)
	require.Equal(t, "", *items[0].Note)
}

func TestUpdateCustomer(t *testing.T) {
	allCustomers, err := repositories.QueryAllCustomers()
	require.NoError(t, err)
	require.Len(t, allCustomers, 4)
	require.Equal(t, "TestInsertCustomer", *allCustomers[3].Name)
	require.Equal(t, "", *allCustomers[3].Note)

	var testCustomerInput = models.NewDefaultCustomerWithNil()
	testCustomerInput.ID = allCustomers[3].ID
	testCustomerInput.Name = models.PtrString("Rea Listik Name")
	testCustomerInput.StreetName = models.PtrString("Parker St")
	testCustomerInput.StreetNumber = models.PtrString("1206")
	testCustomerInput.Note = models.PtrString("complains a lot")

	updatedCustomer, err := repositories.UpdateCustomer(testCustomerInput)
	require.NoError(t, err)
	require.NotNil(t, updatedCustomer)
	require.Len(t, updatedCustomer, 1)
	allCustomers, _ = repositories.QueryAllCustomers()
	require.Equal(t, "Rea Listik Name", *allCustomers[3].Name)
	require.Equal(t, "604-333-3838", *allCustomers[3].Phone)
	require.Equal(t, "Parker St", *allCustomers[3].StreetName)
	require.Equal(t, "1206", *allCustomers[3].StreetNumber)
	require.Equal(t, "", *allCustomers[3].BuzzerNumber)
	require.Equal(t, "complains a lot", *allCustomers[3].Note)
}

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
