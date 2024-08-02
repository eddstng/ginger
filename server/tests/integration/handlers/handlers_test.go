package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"server/db"
	"server/handlers"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
)

var mock pgxmock.PgxConnIface

func TestMain(m *testing.M) {
	var err error
	mock, err = pgxmock.NewConn()
	if err != nil {
		log.Fatalf("Failed to create mock DB connection: %v", err)
	}
	defer mock.Close(context.Background())

	db.SetDBClient(mock)

	// Run the tests
	os.Exit(m.Run())
}

func TestGetItemsHandler(t *testing.T) {
	test_helpers.MockGetItemsQuery(mock)
	handler := handlers.GetItemsHandler()

	req, err := http.NewRequest("GET", "/items", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `[
    {"alcohol":false, "category_id":1, "custom":false, "id":1, "menu_id":1, "name_eng":"Spring Rolls", "name_oth":"春卷", "price":5.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":2, "custom":false, "id":2, "menu_id":2, "name_eng":"Hot and Sour Soup", "name_oth":"酸辣汤", "price":4.99, "special":false, "variant":"Small", "variant_default":true, "variant_price_charge":0},
    {"alcohol":false, "category_id":2, "custom":false, "id":3, "menu_id":2, "name_eng":"Hot and Sour Soup", "name_oth":"酸辣汤", "price":4.99, "special":false, "variant":"Large", "variant_default":false, "variant_price_charge":4},
    {"alcohol":false, "category_id":3, "custom":false, "id":4, "menu_id":3, "name_eng":"Chicken Egg Foo Yung", "name_oth":"雞芙蓉蛋", "price":6.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":4, "custom":false, "id":5, "menu_id":4, "name_eng":"Stir-fried Bok Choy", "name_oth":"炒青菜", "price":7.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":5, "custom":false, "id":6, "menu_id":5, "name_eng":"Salt and Pepper Shrimp", "name_oth":"椒鹽蝦", "price":12.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":6, "custom":false, "id":7, "menu_id":6, "name_eng":"Stir-fried Scallops", "name_oth":"炒帶子", "price":13.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":7, "custom":false, "id":8, "menu_id":7, "name_eng":"Beef Hot Pot", "name_oth":"牛肉煲", "price":14.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":8, "custom":false, "id":9, "menu_id":8, "name_eng":"Sweet and Sour Pork", "name_oth":"糖醋排骨", "price":9.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":9, "custom":false, "id":10, "menu_id":9, "name_eng":"Beef with Broccoli", "name_oth":"西蘭花牛肉", "price":10.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":10, "custom":false, "id":11, "menu_id":10, "name_eng":"Kung Pao Chicken", "name_oth":"宫保鸡丁", "price":8.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":11, "custom":false, "id":12, "menu_id":11, "name_eng":"BBQ Pork Over Rice", "name_oth":"叉燒飯", "price":7.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":12, "custom":false, "id":13, "menu_id":12, "name_eng":"Yangzhou Fried Rice", "name_oth":"扬州炒饭", "price":8.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":13, "custom":false, "id":14, "menu_id":13, "name_eng":"Chicken Chow Mein", "name_oth":"雞肉炒麵", "price":9.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":14, "custom":false, "id":15, "menu_id":14, "name_eng":"Beef Noodle Soup", "name_oth":"牛肉湯麵", "price":6.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":15, "custom":false, "id":16, "menu_id":15, "name_eng":"Pork Congee", "name_oth":"豬肉粥", "price":5.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":false, "category_id":16, "custom":false, "id":17, "menu_id":16, "name_eng":"General Tso's Chicken", "name_oth":"左宗棠雞", "price":15.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0},
    {"alcohol":true, "category_id":17, "custom":false, "id":18, "menu_id":17, "name_eng":"Local Beer", "name_oth":"本地啤酒", "price":2.99, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0}
    ]`
	require.JSONEq(t, expectedBody, rr.Body.String())
}

func TestPostItemHandler(t *testing.T) {
	var testItem = models.NewDefaultItem()
	testItem.NameEng = "TestPostItem"
	testItem.NameOth = "TestPostItemOther"
	testItem.Price = float64(12.50)
	test_helpers.MockInsertItemQuery(mock, testItem)
	handler := handlers.PostItemHandler()

	body, err := createItemJSONRequestBody(testItem)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/items", body)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `[{"alcohol":false, "category_id":1, "custom":false, "id":19, "menu_id":0, "name_eng":"TestPostItem", "name_oth":"TestPostItemOther", "price":12.5, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0}]`
	require.JSONEq(t, expectedBody, rr.Body.String())
}

func TestPutItemHandler(t *testing.T) {
	// testItemInput will be used as the JSON request body.
	var testItemInput = models.NewDefaultItemInput()
	*testItemInput.ID = 18
	*testItemInput.MenuID = 99
	*testItemInput.NameEng = "TestPutItem"
	*testItemInput.NameOth = "TestPutItemOther"
	*testItemInput.Price = float64(12.50)
	*testItemInput.Custom = true
	// repositories.UpdateItem() will recognize the nil and not update the field value.
	testItemInput.Alcohol = nil
	// repositories.UpdateItem() will query the item we are trying to update. Here we mock it.
	test_helpers.MockGetItemQuery(mock, *testItemInput.ID)

	// testItem values will match up with the expected query arguments.
	// The query will be built based on the testitemInput. We use testItem to mock the update query and to set our expected values.
	testItem := models.Item{
		Alcohol:            true,
		CategoryID:         models.GetDefaultCategoryID(),
		Custom:             true,
		ID:                 18,
		MenuID:             99,
		NameEng:            "TestPutItem",
		NameOth:            "TestPutItemOther",
		Price:              12.5,
		Special:            false,
		Variant:            "",
		VariantDefault:     false,
		VariantPriceCharge: 0,
	}

	// Here we mock the update query.
	test_helpers.MockUpdateItemQuery(mock, &testItem)
	handler := handlers.PutItemHandler()

	body, err := createItemInputJSONRequestBody(testItemInput)
	require.NoError(t, err)

	req, err := http.NewRequest("PUT", "/items", body)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	// Expect alcohol to remain true and custom to be false.
	expectedBody := `[{"alcohol":true, "category_id":1, "custom":true, "id":18, "menu_id":99, "name_eng":"TestPutItem", "name_oth":"TestPutItemOther", "price":12.5, "special":false, "variant":"", "variant_default":false, "variant_price_charge":0}]`
	require.JSONEq(t, expectedBody, rr.Body.String())

}

func createItemJSONRequestBody(item *models.Item) (*bytes.Buffer, error) {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonData)
	return body, nil
}

func createItemInputJSONRequestBody(item *models.ItemInput) (*bytes.Buffer, error) {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonData)
	return body, nil
}
