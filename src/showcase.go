package main

import (
	"log"
	"showcase/app"
	"showcase/cmd"
)

func main() {
	server := app.NewServer()
	if err := cmd.Root(server).Execute(); err != nil {
		log.Fatal(err)
	}
}
