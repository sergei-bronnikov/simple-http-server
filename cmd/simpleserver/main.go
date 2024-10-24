package main

import (
	"github.com/sergei-bronnikov/simple-http-server/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
