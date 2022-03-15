package main

import (
	"github.com/mikejeuga/yo_test/src/domain/club"
	"github.com/mikejeuga/yo_test/src/server"
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
