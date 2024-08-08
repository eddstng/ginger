package repositories

import (
	"context"
	"server/db"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryAllCustomersWithMockedDB(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)

	test_helpers.MockGetCustomersQuery(mock)
	items, err := QueryAllCustomers()
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 3)
	require.Equal(t, 1, *items[0].ID)
	expectedItems := []models.Customer{
		{
			ID:           models.PtrInt(1),
			Name:         models.PtrString("John Doe"),
			Phone:        models.PtrString("604-123-1234"),
			UnitNumber:   nil,
			StreetNumber: models.PtrString("5555"),
			StreetName:   models.PtrString("Powel St"),
			BuzzerNumber: nil,
			Note:         nil,
		},
		{
			ID:           models.PtrInt(2),
			Name:         models.PtrString("Christine StClaire"),
			Phone:        models.PtrString("123-456-7890"),
			UnitNumber:   models.PtrString("BSM"),
			StreetNumber: models.PtrString("123"),
			StreetName:   models.PtrString("Maple St"),
			BuzzerNumber: models.PtrString("A12"),
			Note:         models.PtrString("good tips"),
		},
		{
			ID:           models.PtrInt(3),
			Name:         models.PtrString("David Hogan"),
			Phone:        models.PtrString("778-123-1234"),
			UnitNumber:   nil,
			StreetNumber: models.PtrString("5555"),
			StreetName:   models.PtrString("Powel St"),
			BuzzerNumber: nil,
			Note:         nil,
		},
	}
	require.Equal(t, expectedItems[0], items[0])
	require.Equal(t, expectedItems[1], items[1])
	require.Equal(t, expectedItems[2], items[2])
}

func TestInsertCustomerWithMockedDB(t *testing.T) {
	var customer = models.NewDefaultCustomer()
	customer.Name = models.PtrString("John David")
	customer.Phone = models.PtrString("604-341-8384")
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertCustomerQuery(mock, *customer)

	items, err := InsertCustomer(customer)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	expectedItems := []models.Customer{
		{
			ID:           models.PtrInt(19),
			Name:         models.PtrString("John David"),
			Phone:        models.PtrString("604-341-8384"),
			UnitNumber:   models.PtrString(""),
			StreetNumber: models.PtrString(""),
			StreetName:   models.PtrString(""),
			BuzzerNumber: models.PtrString(""),
			Note:         models.PtrString(""),
		},
	}
	require.Equal(t, expectedItems[0], items[0])
}
