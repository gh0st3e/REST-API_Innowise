package service

import (
	"InnowisePreTraineeTask/internal/checks"
	"InnowisePreTraineeTask/internal/entity"
	uuid2 "github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func (s UserService) GetUser(uuid string) (*entity.User, error) {
	s.log.Info("Start Get Userr")
	user, err := s.userRepository.GetUser(uuid)
	if err != nil {
		s.log.Info(err)
		return &entity.User{}, errors.Wrap(err, "service.user.GetUser couldn't get user")
	}

	s.log.Info("End Get User")
	return user, nil
}

func (s UserService) CreateUser(user entity.User) error {
	s.log.Info("Start Create User")

	if user.ID == uuid2.Nil {
		user.ID, _ = uuid2.NewUUID()
	}
	user.Created = time.Now().UTC()

	err := checks.Validation(user)
	if err != nil {
		s.log.Info(err)
		return err
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		s.log.Info(err)
		return errors.Wrap(err, "service.user.CreateUser couldn't create user")
	}

	s.log.Info(user)
	s.log.Info("End Get User")
	return nil
}

func (s UserService) DeleteUser(uuid string) error {
	s.log.Info("Start Delete User")
	err := s.userRepository.DeleteUser(uuid)
	if err != nil {
		s.log.Info(err)
		return errors.Wrap(err, "service.user.DeleteUser couldn't delete user")
	}

	s.log.Info("End Delete User")
	return nil
}

func (s UserService) UpdateUser(uuid string, newUser entity.User) error {
	s.log.Info("Start Update User")
	oldUser, err := s.userRepository.GetUser(uuid)
	if err != nil {
		s.log.Info(err)
		return errors.Wrap(err, "service.user.GetUser couldn't get user")
	}
	s.log.Infof("Old User: %v", oldUser)

	newUser.ID = oldUser.ID
	newUser.Created = oldUser.Created

	err = checks.Validation(newUser)
	if err != nil {
		s.log.Info(err)
		return err
	}

	s.log.Infof("New User: %v", newUser)
	err = s.userRepository.UpdateUser(uuid, newUser)
	if err != nil {
		s.log.Info(err)
		return errors.Wrap(err, "service.user.UpdateUser couldn't update user")
	}

	s.log.Info("End Update User")
	return nil
}

func (s UserService) GetUserList() ([]entity.User, error) {
	s.log.Info("Start Get User List")
	users, err := s.userRepository.GetUserList()
	if err != nil {
		s.log.Info(err)
		return nil, errors.Wrap(err, "service.user.GetUserList couldn't get user list")
	}

	s.log.Info("End Get User List")
	return users, nil
}
