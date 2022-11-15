package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sql.DB {
	//connStr := "user=postgres password=8403 dbname=InnowiseTask sslmode=disable"
	connStr := "postgres://postgres:8403@innowisepretraineetask_db_1/InnowiseTask?sslmode=disable"
	//connStr := "postgres://postgres:8403@172.17.0.5/InnowiseTask?sslmode=disable"
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
