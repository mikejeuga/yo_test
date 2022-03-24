package main

import (
	"github.com/mikejeuga/yo_test/src/internal/adapters/http/server"
	"github.com/mikejeuga/yo_test/src/internal/domain/club"
	"log"
)

func main() {
	tennis := club.NewClub()
	server := server.NewServer(tennis)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Your server has encountered an error at launch time!")
	}
}
