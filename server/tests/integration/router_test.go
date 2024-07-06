package integration_test

import (
	"net/http"
	"net/http/httptest"
	"server/router"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitializeChiRouter(t *testing.T) {
	r := router.InitializeChiRouter()

	t.Run("Root Endpoint", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, "Ginger API", resp.Body.String())
	})
}
