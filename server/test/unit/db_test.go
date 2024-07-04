package unit_test

import (
	"context"
	"server/db"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
)

func TestInitDBClient(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	require.NotNil(t, db.DBClient)
}

func TestCloseDBClient(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	db.SetDBClient(mock)
	mock.ExpectClose()
	err = db.CloseDBClient()
	require.NoError(t, err)
	require.Nil(t, db.DBClient)
}

func TestPgxDBClientQuery(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	db.SetDBClient(mock)
	mock.ExpectQuery("SELECT 1").WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))
	rows, err := db.DBClient.Query(context.Background(), "SELECT 1")
	require.NoError(t, err)
	require.NotNil(t, rows)
}
