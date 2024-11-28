package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/johnsilver94/go-api/cmd/api"
	"github.com/johnsilver94/go-api/configs"
	"github.com/johnsilver94/go-api/db"

	_ "github.com/lib/pq"
)

func main() {
	db, err := db.NewPostgresStorage(configs.Envs.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(fmt.Sprintf(":%s", configs.Envs.Port), db)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)

	}

	log.Println("Connected to database")
}
