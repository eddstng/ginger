package db

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
)

func TestSetupDBClient(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	SetDBClient(mock)
	require.NotNil(t, DBClient)
}

func TestCloseDBClient(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	SetDBClient(mock)
	mock.ExpectClose()
	err = CloseDBClient()
	require.NoError(t, err)
	require.Nil(t, DBClient)
}

func TestPgxDBClientQuery(t *testing.T) {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	defer mock.Close(context.Background())
	SetDBClient(mock)
	mock.ExpectQuery("SELECT 1").WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))
	rows, err := DBClient.Query(context.Background(), "SELECT 1")
	require.NoError(t, err)
	require.NotNil(t, rows)
}
