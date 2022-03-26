package app

import (
	"github.com/mikejeuga/yo_test/src/internal/adapters/http/server"
	"github.com/mikejeuga/yo_test/src/internal/adapters/tennisclubstore"
	"github.com/mikejeuga/yo_test/src/internal/domain/club"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	clubStore := tennisclubstore.InMemoryClubStore{}
	newClub := club.NewClub(&clubStore)

	newServer := server.NewServer(newClub)

	return newServer.ListenAndServe()
}
