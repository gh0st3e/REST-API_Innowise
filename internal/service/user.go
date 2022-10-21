package service

import (
	"InnowisePreTraineeTask/internal/checks"
	"InnowisePreTraineeTask/internal/entity"
	"fmt"
	uuid2 "github.com/google/uuid"
	"time"
)

func (s UserService) GetUser(uuid string) (*entity.User, error) {
	user, err := s.userRepository.GetUser(uuid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (s UserService) CreateUser(user entity.User) error {
	user.ID, _ = uuid2.NewUUID()
	user.Created = time.Now().UTC()

	err := checks.Validation(user)
	if err != nil {
		return err
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	fmt.Println(user)

	return nil
}

func (s UserService) DeleteUser(uuid string) error {
	err := s.userRepository.DeleteUser(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) UpdateUser(uuid string, newUser entity.User) error {
	oldUser, err := s.userRepository.GetUser(uuid)
	if err != nil {
		return err
	}

	newUser.ID = oldUser.ID
	newUser.Created = oldUser.Created

	err = checks.Validation(newUser)
	if err != nil {
		return err
	}

	fmt.Println(newUser)
	err = s.userRepository.UpdateUser(uuid, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) GetUserList() ([]entity.User, error) {
	users, err := s.userRepository.GetUserList()
	if err != nil {
		return nil, err
	}
	return users, nil
}
