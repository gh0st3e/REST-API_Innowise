package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=8403 dbname=InnowiseTask sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Couldn't open database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Couldn't ping database")
	}

	return db
}
