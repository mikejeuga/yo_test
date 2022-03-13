package club_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/yo_test/src/domain/club"
	"github.com/mikejeuga/yo_test/src/models"
	"testing"
)

func TestClubRules(t *testing.T) {
	is := is.New(t)
	t.Run("2 players with the same first and last name cannot be registered into the club", func(t *testing.T) {
		player1 := models.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models.Date{
				Day:   20,
				Month: 7,
				Yeat:  1974,
			},
		}

		player2 := models.Player{
			FirstName:   "Pete",
			LastName:    "Sampras",
			Nationality: "USA",
			DoB: models.Date{
				Day:   15,
				Month: 11,
				Yeat:  1979,
			},
		}

		tennisClub := club.NewClub()
		is.NoErr(tennisClub.Add(player1))
		_ = tennisClub.Add(player2)

		is.Equal(len(tennisClub.Players), 1)
	})
}
