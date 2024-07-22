package main

import (
	"log"

	"github.com/skiba-mateusz/task-manager/cmd/api"
	"github.com/skiba-mateusz/task-manager/config"
	"github.com/skiba-mateusz/task-manager/db"
)

func main() {
	db, err := db.NewPostgreSQLStorage(
		config.Envs.DB.Addr,
		config.Envs.DB.MaxOpenConns,
		config.Envs.DB.MaxIdleConns,
		config.Envs.DB.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("database connection pool established")

	server := api.NewApiServer(config.Envs.Addr, db)

	log.Fatal(server.Run())
}