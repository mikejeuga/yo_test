package club

import (
	"fmt"
	"github.com/mikejeuga/yo_test/src/models"
	"time"
)

type Club struct {
	Players []models.Player
}

func NewClub() *Club {
	var players []models.Player
	return &Club{Players: players}
}

func (c *Club) Add(player models.Player) error {
	if (time.Now().Year() - player.DoB.Year) < 16 || (time.Now().Year() - player.DoB.Year) == 16 && time.Now().Month() > player.DoB.Month  {
		return fmt.Errorf("this player is too young to be part of the club")
	}
	for _, member := range c.Players {
		if member.FirstName == player.FirstName && member.LastName == player.LastName {
			return fmt.Errorf("this player is already registered")
		}
	}
	c.Players = append(c.Players, player)
	return nil
}

