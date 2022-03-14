package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/matryer/is"
	"github.com/mikejeuga/yo_test/src/domain/club"
	"github.com/mikejeuga/yo_test/src/models"
	"github.com/mikejeuga/yo_test/src/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClubServer(t *testing.T) {
	is := is.New(t)
	t.Run("The club server is Healthy", func(t *testing.T) {
		club := club.NewClub()
		clubServer := server.NewServer(club)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		clubServer.Handler.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusOK)
		is.Equal(res.Body.String(), "Welcome to the Tennis Club")
	})

	t.Run("Register a player into the tennis club", func(t *testing.T) {
		tennisClub := club.NewClub()
		server := server.NewServer(tennisClub)

		player := models.Player{
			FirstName:   "Zinedine",
			LastName:    "Zidane",
			Nationality: "FR",
			DoB: models.Date{
				Day:   20,
				Month: 3,
				Year:  1974,
			},
		}

		jsonPlayer, _ := json.Marshal(player)


		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonPlayer))
		res := httptest.NewRecorder()

		server.Handler.ServeHTTP(res,req)

		is.Equal(res.Code, http.StatusCreated)
		is.Equal(tennisClub.Players[0].FirstName, "Zinedine")
		is.Equal(tennisClub.Players[0].Points, models.StartingPoints)

	})
}
