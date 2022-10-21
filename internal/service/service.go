package service

import (
	"InnowisePreTraineeTask/internal/entity"
)

type UserRepository interface {
	GetUser(uuid string) (*entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, user entity.User) error
	GetUserList() ([]entity.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{ur}
}
