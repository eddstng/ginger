package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/db"
	"server/handlers"
	"server/models"
	test_helpers "server/tests/helpers"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
)

func TestGetItemsHandler(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockGetItemsQuery(mock)
	handler := handlers.GetItemsHandler()

	req, err := http.NewRequest("GET", "/items", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `[{"id":1, "category_id": 1, "name_eng":"item1","price":100}]`
	require.JSONEq(t, expectedBody, rr.Body.String())
}

func TestGetItemsHandler_WhenQueryAllItemsFails(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	mock.ExpectQuery("^SELECT (.+) FROM items ORDER BY id ASC$").WillReturnError(errors.New("Mocked Database Error"))
	handler := handlers.GetItemsHandler()

	req, err := http.NewRequest("GET", "/items", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	require.Contains(t, rr.Body.String(), "Mocked Database Error")
}

func TestPostItemHandler(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockInsertItemQuery(mock, models.Item{ID: 111, NameEng: "test_item_insert", Price: 12.50, CategoryID: 0})
	handler := handlers.PostItemHandler()

	item := models.Item{NameEng: "test_item_insert", Price: 12.50, CategoryID: 0}
	body, err := createJSONRequestBody(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/items", body)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `[{"id":111, "category_id": 0, "name_eng":"test_item_insert","price":12.50}]`
	require.JSONEq(t, expectedBody, rr.Body.String())

}

func TestPutItemHandler(t *testing.T) {
	mock, err := test_helpers.SetupPgxMock()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	test_helpers.MockUpdateItemQuery(mock, models.Item{ID: 111, NameEng: "test_item_update", Price: 21.50, CategoryID: 8})
	handler := handlers.PutItemHandler()

	item := models.Item{ID: 111, NameEng: "test_item_update", Price: 21.50, CategoryID: 8}
	body, err := createJSONRequestBody(item)
	require.NoError(t, err)

	req, err := http.NewRequest("PUT", "/items", body)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	expectedBody := `[{"id":111, "category_id": 8, "name_eng":"test_item_update","price":21.50}]`
	require.JSONEq(t, expectedBody, rr.Body.String())

}

func createJSONRequestBody(item models.Item) (*bytes.Buffer, error) {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonData)
	fmt.Println(body.String())
	return body, nil
}
