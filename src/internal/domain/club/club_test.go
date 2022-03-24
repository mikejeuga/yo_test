package club_test

import (
	"github.com/matryer/is"
	models2 "github.com/mikejeuga/yo_test/models"
	"github.com/mikejeuga/yo_test/src/internal/adapters/tennisclubstore"
	"github.com/mikejeuga/yo_test/src/internal/domain/club"
	"testing"
)

func TestClub(t *testing.T) {
	is := is.New(t)
	var testInMemoryClubStore tennisclubstore.InMemoryClubStore
	t.Run("2 players with the same first and last name cannot be registered into the club", func(t *testing.T) {
		player1 := models2.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models2.Date{
				Day:   20,
				Month: 7,
				Year:  1974,
			},
		}

		player2 := models2.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models2.Date{
				Day:   15,
				Month: 11,
				Year:  1979,
			},
		}

		tennisClub := club.NewClub(&testInMemoryClubStore)
		is.NoErr(tennisClub.Add(player1))
		_ = tennisClub.Add(player2)

		is.Equal(len(tennisClub.GetPlayers()), 1)
	})

	t.Run("No player can be less than 16 year's old", func(t *testing.T) {
		player1 := models2.Player{
			FirstName:   "Harry",
			LastName:    "Potter",
			Nationality: "UK",
			DoB: models2.Date{
				Day:   20,
				Month: 10,
				Year:  2012,
			},
		}

		ratpClub := club.NewClub(&testInMemoryClubStore)
		err := ratpClub.Add(player1)

		is.Equal(err.Error(), "this player is too young to be part of the club")
	})

	t.Run("registered player should start with the score of 1200 points", func(t *testing.T) {
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

		wtaClub := club.NewClub(&testInMemoryClubStore)
		err := wtaClub.Add(player)

		is.NoErr(err)
		is.Equal(wtaClub.GetPlayers()[0].Points, models2.StartingPoints)
	})
}
