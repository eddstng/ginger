package repositories_test

import (
	"server/models"
	"server/repositories"
	"testing"

	"github.com/stretchr/testify/require"
)

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
