package main

import (
	"log"

	"github.com/skiba-mateusz/task-manager/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080")


	log.Fatal(server.Run())
}