package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=8403 dbname=InnowiseTask sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
