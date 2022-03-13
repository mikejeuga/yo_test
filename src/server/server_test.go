package server_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/yo_test/src/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClubServer(t *testing.T) {
	is := is.New(t)
	t.Run("The club server is Healthy", func(t *testing.T) {
		clubServer := server.NewServer()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		clubServer.Handler.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)
		is.Equal(res.Body.String(), "Welcome to the Tennis Club")
	})
}
