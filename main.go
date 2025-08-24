package main

import (
	"log"

	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("There is no connection to DB")
		return
	}

	handlers.Handlers()
}
