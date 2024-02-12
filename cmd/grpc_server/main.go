package main

import (
	"log"

	"github.com/terdenan/msc_chat_server/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to instantiate app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
