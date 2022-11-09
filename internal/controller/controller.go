package controller

import (
	"InnowisePreTraineeTask/internal/entity"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetUser(uuid string) (*entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, newUser entity.User) error
	GetUserList() ([]entity.User, error)
}

type UserController struct {
	log         *logrus.Logger
	userService UserService
}

func NewUserServer(log *logrus.Logger, us UserService) *UserController {
	return &UserController{log, us}
}
