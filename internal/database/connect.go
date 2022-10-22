package database

import (
	"InnowisePreTraineeTask/internal/util"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", util.UserDB, util.PasswordDB, util.NameDB, util.SslMode)
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
