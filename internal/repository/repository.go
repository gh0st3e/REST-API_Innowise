package repository

import (
	"InnowisePreTraineeTask/internal/entity"
	"database/sql"
)

type IRepository interface {
	GetUser(uuid string) (*entity.User, error)
	//CreateUser(user entity.User) error
	//DeleteUser(uuid string) error
	//UpdateUser(uuid string, user entity.User) error
	//GetUserList() ([]entity.User, error)
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) IRepository {
	return Repo{db}
}
