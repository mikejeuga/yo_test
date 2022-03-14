package club_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/yo_test/src/domain/club"
	"github.com/mikejeuga/yo_test/src/models"
	"testing"
)

func TestClub(t *testing.T) {
	is := is.New(t)
	t.Run("2 players with the same first and last name cannot be registered into the club", func(t *testing.T) {
		player1 := models.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models.Date{
				Day:   20,
				Month: 7,
				Year:  1974,
			},
		}

		player2 := models.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models.Date{
				Day:   15,
				Month: 11,
				Year:  1979,
			},
		}

		tennisClub := club.NewClub()
		is.NoErr(tennisClub.Add(player1))
		_ = tennisClub.Add(player2)

		is.Equal(len(tennisClub.Players), 1)
	})

	t.Run("No player can be less than 16 year's old", func(t *testing.T) {
		player1 := models.Player{
			FirstName:   "Harry",
			LastName:    "Potter",
			Nationality: "UK",
			DoB: models.Date{
				Day:   20,
				Month: 10,
				Year:  2012,
			},
		}

		ratpClub := club.NewClub()
		err := ratpClub.Add(player1)

		is.Equal(err.Error(), "this player is too young to be part of the club")
	})

	t.Run("registered player should start with the score of 1200 points", func(t *testing.T) {
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
		wtaClub := club.NewClub()
		err := wtaClub.Add(player)

		is.NoErr(err)
		is.Equal(wtaClub.Players[0].Points, models.StartingPoints)
	})
}
