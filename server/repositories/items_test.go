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
	require.Len(t, items, 18)
	require.Equal(t, 1, items[0].ID)
	require.Equal(t, "Spring Rolls", items[0].NameEng)
	require.Equal(t, float64(5.99), items[0].Price)
	expectedItems := []models.Item{
		{
			ID:                 1,
			MenuID:             1,
			CategoryID:         1,
			Price:              5.99,
			NameEng:            "Spring Rolls",
			NameOth:            "春卷",
			Special:            false,
			Alcohol:            false,
			Custom:             false,
			Variant:            "",
			VariantDefault:     false,
			VariantPriceCharge: 0,
		},
		{
			ID:                 2,
			MenuID:             2,
			CategoryID:         2,
			Price:              4.99,
			NameEng:            "Hot and Sour Soup",
			NameOth:            "酸辣汤",
			Special:            false,
			Alcohol:            false,
			Custom:             false,
			Variant:            "Small",
			VariantDefault:     true,
			VariantPriceCharge: 0,
		},
		{
			ID:                 3,
			MenuID:             2,
			CategoryID:         2,
			Price:              4.99,
			NameEng:            "Hot and Sour Soup",
			NameOth:            "酸辣汤",
			Special:            false,
			Alcohol:            false,
			Custom:             false,
			Variant:            "Large",
			VariantDefault:     false,
			VariantPriceCharge: 4.00,
		},
	}
	require.Equal(t, expectedItems[0], items[0])
	require.Equal(t, expectedItems[1], items[1])
	require.Equal(t, expectedItems[2], items[2])
}

func TestInsertItemWithMockedDB(t *testing.T) {
	var testItem = models.NewDefaultItem()
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertItemQuery(mock, testItem)
	items, err := InsertItem(testItem)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)

	expectedItems := []models.Item{
		{
			ID:                 19,
			MenuID:             0,
			CategoryID:         1, // Default Category ID
			Price:              0.00,
			NameEng:            "",
			NameOth:            "",
			Special:            false,
			Alcohol:            false,
			Custom:             false,
			Variant:            "",
			VariantDefault:     false,
			VariantPriceCharge: 0,
		},
	}
	require.Equal(t, expectedItems[0], items[0])
}

func TestUpdateItemWithMockedDB(t *testing.T) {
	// The testItemInput will only contain values for fields that are being updated. The rest of the values will be nil in the testItemInput.
	// The nil items will be handled by the updateItemFields function in the repositories/items.go file and will not be updated.
	var testItemInput = models.NewDefaultItemInputWithNil()
	testItemInput.ID = new(int)
	testItemInput.NameEng = new(string)
	testItemInput.Special = new(bool)
	*testItemInput.ID = 18
	*testItemInput.NameEng = "Local Beer - Updated"
	*testItemInput.Special = true

	var expectedTestedItem = models.NewDefaultItem()
	expectedTestedItem.ID = 18
	expectedTestedItem.MenuID = 17
	expectedTestedItem.CategoryID = 17
	expectedTestedItem.Price = 2.99
	expectedTestedItem.NameEng = "Local Beer - Updated"
	expectedTestedItem.NameOth = "本地啤酒"
	expectedTestedItem.Special = true
	expectedTestedItem.Alcohol = true
	expectedTestedItem.Custom = false
	expectedTestedItem.Variant = ""
	expectedTestedItem.VariantDefault = false
	expectedTestedItem.VariantPriceCharge = 0.00

	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockGetItemQuery(mock, *testItemInput.ID)
	test_helpers.MockUpdateItemQuery(mock, expectedTestedItem)
	items, err := UpdateItem(testItemInput)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, items[0].ID, *testItemInput.ID)
	require.Equal(t, items[0].NameEng, *testItemInput.NameEng)
	// Item 18 started as false and was modified to true in testItemInput, so it should be true.
	require.Equal(t, items[0].Special, true)
	// Item 18 started as true and was not modified in testItemInput, so it should still be true.
	require.Equal(t, items[0].Alcohol, true)
}
