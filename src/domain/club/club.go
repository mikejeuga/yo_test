package club

import (
	"fmt"
	"github.com/mikejeuga/yo_test/src/models"
)

type Club struct {
	Players []models.Player
}

func NewClub() *Club {
	var players []models.Player
	return &Club{Players: players}
}

func (c *Club) Add(player models.Player) error {
	for _, member := range c.Players {
		if member.FirstName == player.FirstName && member.LastName == player.LastName {
			return fmt.Errorf("this player is already registered")
		}
	}
	c.Players = append(c.Players, player)
	return nil
}

