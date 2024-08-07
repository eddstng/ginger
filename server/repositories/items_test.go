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
	require.Equal(t, 1, *items[0].ID)
	require.Equal(t, "Spring Rolls", *items[0].NameEng)
	require.Equal(t, float64(5.99), *items[0].Price)
	expectedItems := []models.Item{
		{
			ID:                 models.PtrInt(1),
			MenuID:             models.PtrInt(1),
			CategoryID:         models.PtrInt(1),
			Price:              models.PtrFloat64(5.99),
			NameEng:            models.PtrString("Spring Rolls"),
			NameOth:            models.PtrString("春卷"),
			Special:            models.PtrBool(false),
			Alcohol:            models.PtrBool(false),
			Custom:             models.PtrBool(false),
			Variant:            models.PtrString(""),
			VariantDefault:     models.PtrBool(false),
			VariantPriceCharge: models.PtrFloat64(0),
		},
		{
			ID:                 models.PtrInt(2),
			MenuID:             models.PtrInt(2),
			CategoryID:         models.PtrInt(2),
			Price:              models.PtrFloat64(4.99),
			NameEng:            models.PtrString("Hot and Sour Soup"),
			NameOth:            models.PtrString("酸辣汤"),
			Special:            models.PtrBool(false),
			Alcohol:            models.PtrBool(false),
			Custom:             models.PtrBool(false),
			Variant:            models.PtrString("Small"),
			VariantDefault:     models.PtrBool(true),
			VariantPriceCharge: models.PtrFloat64(0),
		},
		{
			ID:                 models.PtrInt(3),
			MenuID:             models.PtrInt(2),
			CategoryID:         models.PtrInt(2),
			Price:              models.PtrFloat64(4.99),
			NameEng:            models.PtrString("Hot and Sour Soup"),
			NameOth:            models.PtrString("酸辣汤"),
			Special:            models.PtrBool(false),
			Alcohol:            models.PtrBool(false),
			Custom:             models.PtrBool(false),
			Variant:            models.PtrString("Large"),
			VariantDefault:     models.PtrBool(false),
			VariantPriceCharge: models.PtrFloat64(4.00),
		},
	}
	require.Equal(t, expectedItems[0], items[0])
	require.Equal(t, expectedItems[1], items[1])
	require.Equal(t, expectedItems[2], items[2])
}

func TestInsertItemWithMockedDB(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.MenuID = models.PtrInt(99)
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertItemQuery(mock, *testItem)

	items, err := InsertItem(testItem)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	expectedItems := []models.Item{
		{
			ID:                 models.PtrInt(19),
			MenuID:             models.PtrInt(99),
			CategoryID:         models.PtrInt(1), // Default Category ID
			Price:              models.PtrFloat64(0.00),
			NameEng:            models.PtrString(""),
			NameOth:            models.PtrString(""),
			Special:            models.PtrBool(false),
			Alcohol:            models.PtrBool(false),
			Custom:             models.PtrBool(false),
			Variant:            models.PtrString(""),
			VariantDefault:     models.PtrBool(false),
			VariantPriceCharge: models.PtrFloat64(0),
		},
	}
	require.Equal(t, expectedItems[0], items[0])
}

func TestUpdateItemWithMockedDB(t *testing.T) {
	// The testItem will only contain values for fields that are being updated. The rest of the values will be nil in the testItem.
	// The nil items will be handled by the updateItemFields function in the repositories/items.go file and will not be updated.
	var testItem = models.NewDefaultItemWithNil()
	testItem.ID = models.PtrInt(18)
	testItem.NameEng = models.PtrString("Local Beer - Updated")
	testItem.Special = models.PtrBool(true)

	var expectedTestedItem = models.NewDefaultItem()
	expectedTestedItem.ID = models.PtrInt(18)
	expectedTestedItem.MenuID = models.PtrInt(17)
	expectedTestedItem.CategoryID = models.PtrInt(17)
	expectedTestedItem.Price = models.PtrFloat64(2.99)
	expectedTestedItem.NameEng = models.PtrString("Local Beer - Updated")
	expectedTestedItem.NameOth = models.PtrString("本地啤酒")
	expectedTestedItem.Special = models.PtrBool(true)
	expectedTestedItem.Alcohol = models.PtrBool(true)
	expectedTestedItem.Custom = models.PtrBool(false)
	expectedTestedItem.Variant = models.PtrString("")
	expectedTestedItem.VariantDefault = models.PtrBool(false)
	expectedTestedItem.VariantPriceCharge = models.PtrFloat64(0.00)

	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockGetItemQuery(mock, *testItem.ID)
	test_helpers.MockUpdateItemQuery(mock, *expectedTestedItem)
	items, err := UpdateItem(testItem)
	require.NoError(t, err)
	require.NotNil(t, items)
	require.Len(t, items, 1)
	require.Equal(t, *items[0].ID, *testItem.ID)
	require.Equal(t, *items[0].NameEng, *testItem.NameEng)
	// Item 18 started as false and was modified to true in testItem, so it should be true.
	require.Equal(t, *items[0].Special, true)
	// Item 18 started as true and was not modified in testItem, so it should still be true.
	require.Equal(t, *items[0].Alcohol, true)
}
