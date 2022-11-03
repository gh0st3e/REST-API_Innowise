package repository

import (
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/service"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"regexp"
	"testing"
	"time"
)

func Connect() *service.UserService {
	connStr := "user=postgres password=8403 dbname=InnowiseTaskTest sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Couldn't open database + %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Couldn't ping database")
	}

	log := logrus.New()
	userRepository := NewUserRepository(db)
	userService := service.NewUserService(log, userRepository)

	return userService
}

func TestCreateUserWithMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("err: %s", err)
	}
	defer db.Close()

	r := NewUserRepository(db)
	log := logrus.New()
	s := service.NewUserService(log, r)

	type mockBehaviour func(user entity.User, id int)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		testUser      entity.User
		id            int
		expected      error
	}{
		{
			name: "Test Add One",
			testUser: entity.User{
				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
				Created:   time.Now(),
			},
			id: 2,
			mockBehaviour: func(user entity.User, id int) {
				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("INSERT INTO users VALUES('%s','%s','%s','%s',%v,'%v')", user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created.Format(time.RFC3339)))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	//"INSERT INTO users VALUES('18e90062-54b0-11ed-86a7-e8d8d1f76e0b','Viktor','Korneplod','viktor.korneplod.2020@mail.ru',100,'2022-11-03T03:34:02+03:00')"
	//"INSERT INTO users VALUES('18e90062-54b0-11ed-86a7-e8d8d1f76e0b','Viktor','Korneplod','viktor.korneplod.2020@mail.ru',100,'2022-11-03T03:34:02+03:00')"
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.mockBehaviour(testCase.testUser, testCase.id)

			actual := service.UserRepository.CreateUser(s, testCase.testUser)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

//func TestCreateUser(t *testing.T) {
//
//	userService := Connect()
//
//	testTable := []struct {
//		name     string
//		testUser entity.User
//		expected error
//	}{
//		{
//			name: "Test Add One",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "Viktor",
//				Lastname:  "Korneplod",
//				Email:     "viktor.korneplod.2020@mail.ru",
//				Age:       100,
//			},
//			expected: nil,
//		},
//		{
//			name: "Test Add Two",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "Fast",
//				Lastname:  "Ganzales",
//				Email:     "fast.ganzales.2003@yandex.ru",
//				Age:       100,
//			},
//			expected: nil,
//		},
//		{
//			name: "Test Add Three",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18eaabba-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "George",
//				Lastname:  "Floyd",
//				Email:     "george.floyd.2003@breathe.en",
//				Age:       100,
//			},
//			expected: nil,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			actual := service.UserRepository.CreateUser(userService, testCase.testUser)
//
//			if actual != testCase.expected {
//				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
//			}
//		})
//	}
//}
//
//func TestUpdateUser(t *testing.T) {
//	userService := Connect()
//	testTable := []struct {
//		name     string
//		testUser entity.User
//		expected error
//	}{
//		{
//			name: "Test Update One",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "Denis",
//				Lastname:  "Leonov",
//				Email:     "denis.leonov.2003@yandex.ru",
//				Age:       19,
//			},
//			expected: nil,
//		},
//		{
//			name: "Test Update Two",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "Andrey",
//				Lastname:  "Ivanov",
//				Email:     "andrey.ivanov.2003@yandex.ru",
//				Age:       19,
//			},
//			expected: nil,
//		},
//		{
//			name: "Test Update Three",
//			testUser: entity.User{
//				ID:        uuid.MustParse("18eaabba-54b0-11ed-86a7-e8d8d1f76e0b"),
//				Firstname: "Vladimir",
//				Lastname:  "Putin",
//				Email:     "vladimir.putin.2022@russia.ru",
//				Age:       60,
//			},
//			expected: nil,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			actual := service.UserRepository.UpdateUser(userService, testCase.testUser.ID.String(), testCase.testUser)
//
//			if actual != testCase.expected {
//				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
//			}
//		})
//	}
//}
//
//func TestDeleteUser(t *testing.T) {
//	userService := Connect()
//	testTable := []struct {
//		name     string
//		id       string
//		expected error
//	}{
//		{
//			name:     "Test Delete One",
//			id:       "18e90062-54b0-11ed-86a7-e8d8d1f76e0b",
//			expected: nil,
//		},
//		{
//			name:     "Test Delete Two",
//			id:       "18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b",
//			expected: nil,
//		},
//		{
//			name:     "Test Delete Three",
//			id:       "18eaabba-54b0-11ed-86a7-e8d8d1f76e0b",
//			expected: nil,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			actual := service.UserRepository.DeleteUser(userService, testCase.id)
//
//			if actual != testCase.expected {
//				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
//			}
//		})
//	}
//}
