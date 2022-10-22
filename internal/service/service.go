package service

import (
	"InnowisePreTraineeTask/internal/entity"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	GetUser(uuid string) (*entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, user entity.User) error
	GetUserList() ([]entity.User, error)
}

type UserService struct {
	log            *logrus.Logger
	userRepository UserRepository
}

func NewUserService(log *logrus.Logger, ur UserRepository) *UserService {
	return &UserService{log, ur}
}
