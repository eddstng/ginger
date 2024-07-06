package integration_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/db"
	"server/handlers"
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
	mock.ExpectQuery("^SELECT (.+) FROM items$").WillReturnError(errors.New("Mocked Database Error"))
	handler := handlers.GetItemsHandler()

	req, err := http.NewRequest("GET", "/items", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	require.Contains(t, rr.Body.String(), "Mocked Database Error")
}
