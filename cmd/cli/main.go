package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("postgres", "user=kseikseich password=11 dbname=yap sslmode=disable")
	if err != nil {
		return
	}
	// setup database

	if err := goose.Up(db, "../../db/migrations"); err != nil {
		panic(err)
	}

	// run app
}
