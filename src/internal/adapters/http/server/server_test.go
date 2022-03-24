package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/matryer/is"
	models2 "github.com/mikejeuga/yo_test/models"
	"github.com/mikejeuga/yo_test/src/internal/adapters/http/server"
	"github.com/mikejeuga/yo_test/src/internal/adapters/tennisclubstore"
	"github.com/mikejeuga/yo_test/src/internal/domain/club"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClubServer(t *testing.T) {
	is := is.New(t)
	var testInMemoryClubStore tennisclubstore.InMemoryClubStore
	t.Run("The club server is Healthy", func(t *testing.T) {
		club := club.NewClub(&testInMemoryClubStore)
		clubServer := server.NewServer(club)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		clubServer.Handler.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)
		is.Equal(res.Body.String(), "Welcome to the Tennis Club")
	})

	t.Run("Register a player into the tennis club", func(t *testing.T) {
		tennisClub := club.NewClub(&testInMemoryClubStore)
		server := server.NewServer(tennisClub)

		player := models2.Player{
			FirstName:   "Zinedine",
			LastName:    "Zidane",
			Nationality: "FR",
			DoB: models2.Date{
				Day:   20,
				Month: 3,
				Year:  1974,
			},
		}

		jsonPlayer, _ := json.Marshal(player)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonPlayer))
		res := httptest.NewRecorder()

		server.Handler.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusCreated)
		is.Equal(tennisClub.GetPlayers()[0].FirstName, "Zinedine")
		is.Equal(tennisClub.GetPlayers()[0].Points, models2.StartingPoints)

	})
}
