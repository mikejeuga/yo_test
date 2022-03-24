package tennisclubstore

import (
	"fmt"
	"github.com/mikejeuga/yo_test/models"
)

type InMemoryClubStore []models.Player

func (i *InMemoryClubStore) GetPlayers() []models.Player {
	return *i
}

func (i *InMemoryClubStore) AddPlayer(player models.Player) {
	*i = append(*i, player)
}

func (i InMemoryClubStore) IsPlayerRegistered(player models.Player) error {
	for _, member := range i {
		if member.FirstName == player.FirstName && member.LastName == player.LastName {
			return fmt.Errorf("this player is already registered")
		}
	}
	return nil
}
