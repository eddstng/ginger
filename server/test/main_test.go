package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"server/db"
	"server/router"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func executeRequest(req *http.Request, router *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestChiRouter(t *testing.T) {
	s := router.InitializeChiRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "Ginger API", response.Body.String())
}

func TestDBClient(t *testing.T) {
	err := godotenv.Load("./.env")
	require.Nil(t, err)
	err = db.InitDBClient(os.Getenv("DATABASE_URL"))
	require.Nil(t, err)
	rows, err := db.DBClient.Query(context.Background(), "SELECT id, name_eng, price, category_id FROM items LIMIT 1")
	require.Nil(t, err)
	require.NotNil(t, db.DBClient)
	var item db.Item
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.NameEng, &item.Price, &item.CategoryID)
	}
	require.Nil(t, err)
	require.Equal(t, 1, item.ID)
	db.CloseDBClient()
	require.Nil(t, db.DBClient)
}
