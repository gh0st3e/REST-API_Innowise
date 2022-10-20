package repository

import (
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/entity"
	"fmt"
)

var queryUser = map[string]string{
	"CreateUser":  "INSERT INTO users VALUES('%s','%s','%s','%s',%v,'%v')",
	"GetUser":     "SELECT * FROM users WHERE id = '%s'",
	"DeleteUser":  "DELETE FROM users WHERE id='%s'",
	"UpdateUser":  "UPDATE users SET firstname='%s',lastname='%s',email='%s',age=%v WHERE id='%s'",
	"GetUserList": "SELECT * FROM users",
}

func (r Repo) GetUser(uuid string) (*entity.User, error) {
	query := fmt.Sprintf(queryUser["GetUser"], uuid)
	result, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	result.Next()
	user := entity.User{}
	err = result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func CreateUser(user entity.User) error {
	db := database.Connect()
	//TODO переделать запрос
	query := fmt.Sprintf(queryUser["CreateUser"], user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created.Format("2006.01.02 15:04:05"))
	result, err := db.Exec(query)
	if err != nil {
		_ = result
		panic(err)
		//return err
	}
	return nil
}

func DeleteUser(uuid string) error {
	db := database.Connect()
	query := fmt.Sprintf(queryUser["DeleteUser"], uuid)
	result, err := db.Exec(query)
	if err != nil {
		_ = result
		panic(err)
	}
	return nil
}

func UpdateUser(uuid string, user entity.User) error {
	db := database.Connect()
	query := fmt.Sprintf(queryUser["UpdateUser"], user.Firstname, user.Lastname, user.Email, user.Age, uuid)
	result, err := db.Exec(query)
	if err != nil {
		_ = result
		return err
	}
	return nil
}

func GetUserList() ([]entity.User, error) {
	db := database.Connect()
	query := fmt.Sprintf(queryUser["GetUserList"])
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	users := []entity.User{}

	for result.Next() {
		user := entity.User{}
		err = result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
