package repository

import (
	"fmt"
)

func (r UserRepository) CheckTable() error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS users" +
		"(\n    id UUID NOT NULL,\n" +
		"    CONSTRAINT id_user PRIMARY KEY(id),\n " +
		"   firstname CHARACTER VARYING(30),\n" +
		"    lastname CHARACTER VARYING(30),\n" +
		"    email CHARACTER VARYING(40) UNIQUE,\n" +
		"    age integer,\n" +
		"    created timestamp\n);")

	_, err := r.db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
