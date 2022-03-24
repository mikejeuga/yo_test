package tennisclubstore

import (
	"github.com/mikejeuga/yo_test/models"
)

type InMemoryClubStore []models.Player

func (i *InMemoryClubStore) GetPlayers() []models.Player {
	return *i
}

func (i *InMemoryClubStore) AddPlayer(player models.Player) {
	*i = append(*i, player)
}
