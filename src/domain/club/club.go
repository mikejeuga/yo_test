package club

import (
	"fmt"
	"github.com/mikejeuga/yo_test/src/models"
	"time"
)

type Club struct {
	Players []models.Player
}

type TennisClub interface {
	Add(player models.Player) error
}


func NewClub() *Club {
	var players []models.Player
	return &Club{Players: players}
}

func (c *Club) Add(player models.Player) error {
	err := checkAge(player)
	if err != nil {
		return err
	}

	err = c.checkIdentity(player)
	if err != nil {
		return err
	}

	player.Score(models.StartingPoints)

	c.Players = append(c.Players, player)
	return nil
}

func checkAge(player models.Player) error {
	if (time.Now().Year() - player.DoB.Year) < 16 || (time.Now().Year() - player.DoB.Year) == 16 && time.Now().Month() > player.DoB.Month  {
		return fmt.Errorf("this player is too young to be part of the club")
	}
	return nil
}

func (c *Club)checkIdentity(player models.Player) error {
	for _, member := range c.Players {
		if member.FirstName == player.FirstName && member.LastName == player.LastName {
			return fmt.Errorf("this player is already registered")
		}
	}
	return nil
}
