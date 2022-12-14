package repository

import (
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/service"
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

func TestCreateUserWithMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("err: %s", err)
	}
	defer db.Close()

	r := NewUserRepository(db)

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
			id: 1,
			mockBehaviour: func(user entity.User, id int) {

				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("INSERT INTO \"users\" VALUES('%s','%s','%s','%s',%v,'%v')", user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created.Format(time.RFC3339)))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name: "Test Add Two",
			testUser: entity.User{
				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Boris",
				Lastname:  "Johnson",
				Email:     "boris.johnson.2022@mail.ru",
				Age:       45,
				Created:   time.Now(),
			},
			id: 2,
			mockBehaviour: func(user entity.User, id int) {

				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("INSERT INTO \"users\" VALUES('%s','%s','%s','%s',%v,'%v')", user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created.Format(time.RFC3339)))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.mockBehaviour(testCase.testUser, testCase.id)

			actual := UserRepository.CreateUser(*r, testCase.testUser)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestUpdateUserWithMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("err: %s", err)
	}
	defer db.Close()

	r := NewUserRepository(db)

	type mockBehaviour func(user entity.User, id int)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		testUser      entity.User
		id            int
		expected      error
	}{
		{
			name: "Test Update One",
			testUser: entity.User{
				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "upd_Viktor",
				Lastname:  "upd_Kornep",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
				Created:   time.Now(),
			},
			id: 1,
			mockBehaviour: func(user entity.User, id int) {
				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("UPDATE \"users\" SET firstname='%s',lastname='%s',email='%s',age=%v WHERE id='%s'", user.Firstname, user.Lastname, user.Email, user.Age, user.ID))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name: "Test Update Two",
			testUser: entity.User{
				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "upd_Boris",
				Lastname:  "upd_John",
				Email:     "boris.johnson.2022@mail.ru",
				Age:       45,
				Created:   time.Now(),
			},
			id: 2,
			mockBehaviour: func(user entity.User, id int) {
				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("UPDATE \"users\" SET firstname='%s',lastname='%s',email='%s',age=%v WHERE id='%s'", user.Firstname, user.Lastname, user.Email, user.Age, user.ID))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.mockBehaviour(testCase.testUser, testCase.id)

			actual := UserRepository.UpdateUser(*r, testCase.testUser.ID.String(), testCase.testUser)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestDeleteUserWithMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("err: %s", err)
	}
	defer db.Close()

	r := NewUserRepository(db)
	log := logrus.New()
	s := service.NewUserService(log, r)

	type mockBehaviour func(userID string, id int)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		userID        string
		id            int
		expected      error
	}{
		{
			name:   "Test Delete One",
			userID: "18e90062-54b0-11ed-86a7-e8d8d1f76e0b",
			id:     1,
			mockBehaviour: func(userID string, id int) {
				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("DELETE FROM \"users\" WHERE id='%s'", userID))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name:   "Test Delete Two",
			userID: "18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b",
			id:     2,
			mockBehaviour: func(userID string, id int) {
				mock.ExpectExec(regexp.QuoteMeta(fmt.Sprintf("DELETE FROM \"users\" WHERE id='%s'", userID))).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.mockBehaviour(testCase.userID, testCase.id)

			actual := service.UserRepository.DeleteUser(s, testCase.userID)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

//func Connect() *service.UserService {
//	connStr := "user=postgres password=8403 dbname=InnowiseTaskTest sslmode=disable"
//
//	db, err := sql.Open("postgres", connStr)
//
//	if err != nil {
//		log.Fatalf("Couldn't open database + %s", err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		log.Fatal("Couldn't ping database")
//	}
//
//	log := logrus.New()
//	userRepository := NewUserRepository(db)
//	userService := service.NewUserService(log, userRepository)
//
//	return userService
//}

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
