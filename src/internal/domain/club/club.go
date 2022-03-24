package club

import (
	"fmt"
	"github.com/mikejeuga/yo_test/models"
	"time"
)

type TennisClub interface {
	Add(player models.Player) error
}

type Club struct {
	PlayerStore TennisClubStore
}

type TennisClubStore interface {
	AddPlayer(player models.Player)
	GetPlayers() []models.Player
}

func NewClub(clubStore TennisClubStore) *Club {
	return &Club{PlayerStore: clubStore}
}

func (c *Club) Add(player models.Player) error {
	err := checkAge(player)
	if err != nil {
		return err
	}

	if c.hasRegistered(player) {
		return fmt.Errorf("This player is already a member of the club")
	}

	player.Score(models.StartingPoints)

	c.PlayerStore.AddPlayer(player)
	return nil
}

func (c *Club) GetPlayers() []models.Player {
	return c.PlayerStore.GetPlayers()
}

func (c *Club) hasRegistered(player models.Player) bool {
	for _, member := range c.PlayerStore.GetPlayers() {
		if member.FirstName == player.FirstName && member.LastName == player.LastName {
			return true
		}
	}
	return false
}

func checkAge(player models.Player) error {
	if (time.Now().Year()-player.DoB.Year) < 16 || (time.Now().Year()-player.DoB.Year) == 16 && time.Now().Month() > player.DoB.Month {
		return fmt.Errorf("this player is too young to be part of the club")
	}
	return nil
}
