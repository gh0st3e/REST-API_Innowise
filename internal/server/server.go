package server

import (
	"InnowisePreTraineeTask/internal/entity"
)

type UserService interface {
	GetUser(uuid string) (*entity.User, error)
	CreateUser(user entity.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, newUser entity.User) error
	GetUserList() ([]entity.User, error)
}

type UserServer struct {
	userService UserService
}

func NewUserServer(us UserService) *UserServer {
	return &UserServer{us}
}
