package main

import (
	"github.com/mikejeuga/yo_test/src/app"
	"log"
)

func main() {
	newApp := app.NewApp()
	err := newApp.Start()
	if err != nil {
		log.Fatal(err)
	}
}
