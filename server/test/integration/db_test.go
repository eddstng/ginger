package integration

import (
	"os"
	"server/db"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestInitDBClientFromURL(t *testing.T) {
	err := godotenv.Load("../../.env")
	databaseURL := os.Getenv("DATABASE_URL")
	require.NoError(t, err)
	err = db.InitDBClientFromURL(databaseURL)
	require.NoError(t, err)
	require.NotNil(t, db.DBClient)
}
