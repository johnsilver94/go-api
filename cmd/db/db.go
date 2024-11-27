package db

import (
	"database/sql"
	"log"
)

func NewPostgresStorage(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
