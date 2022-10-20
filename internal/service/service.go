package service

import (
	"InnowisePreTraineeTask/internal/repository"
)

type IService interface {
	GetUser(uuid string) ([]byte, error)
	//CreateUser(user entity.User) error
	//DeleteUser(uuid string) error
	//UpdateUser(uuid string, user entity.User) error
	//GetUserList() ([]entity.User, error)
}

type Store struct {
	rp repository.IRepository
}

func New(repo repository.IRepository) IService {
	return &Store{repo}
}
