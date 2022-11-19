package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sql.DB {
	//connStr := "user=postgres password=8403 dbname=InnowiseTask sslmode=disable" // Local DB
	connStr := "postgres://postgres:8403@db:5432/InnowiseTask?sslmode=disable" // Minikube DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Couldn't open database")
	}

	err = db.Ping()
	if err != nil {

		log.Fatalf("Couldn't ping database. Err:%s", err)
	}

	return db
}
