package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=root host=localhost dbname=slc3 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}
	return db
}