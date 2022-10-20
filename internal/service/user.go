package service

import (
	"InnowisePreTraineeTask/internal/checks"
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/repository"
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"io"
	"time"
)

func (s *Store) GetUser(uuid string) ([]byte, error) {
	user, err := s.rp.GetUser(uuid)
	//user, err := repository.Repo.GetUser(r,uuid)
	//user, err := repository.GetUser(uuid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byteUser, nil
}

func CreateUser(closer io.ReadCloser) error {
	var user entity.User
	_ = json.NewDecoder(closer).Decode(&user)
	user.ID, _ = uuid2.NewUUID()
	user.Created = time.Now().UTC()

	err := checks.Validation(user)
	if err != nil {
		return err
	}

	repository.CreateUser(user)
	fmt.Println(user)

	return nil
}

func DeleteUser(uuid string) error {
	err := repository.DeleteUser(uuid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(uuid string, closer io.ReadCloser) error {
	//oldUser, err := repository.Repo{}.GetUser(uuid)
	////oldUser, err := repository.GetUser(uuid)
	//if err != nil {
	//	return err
	//}
	//
	//var newUser entity.User
	//_ = json.NewDecoder(closer).Decode(&newUser)
	//
	//newUser.ID = oldUser.ID
	//newUser.Created = oldUser.Created
	//
	//err = checks.Validation(newUser)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(newUser)
	//err = repository.UpdateUser(uuid, newUser)
	//if err != nil {
	//	return err
	//}
	//
	return nil
}

func GetUserList() ([]byte, error) {
	users, err := repository.GetUserList()
	if err != nil {
		return nil, err
	}
	byteUsers, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return byteUsers, nil
}
