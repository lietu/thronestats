package main

import (
	"log"
	"github.com/lietu/thronestats/thronestats"
)

func main() {
	settings := thronestats.GetServerSettings()
	thronestats.RunServer(settings)
	log.Fatalf("Exiting...")
}
