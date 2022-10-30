package tests

import (
	"InnowisePreTraineeTask/internal/checks"
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/repository"
	"InnowisePreTraineeTask/internal/service"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"testing"
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
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(log, userRepository)

	return userService
}

func TestCreateUser(t *testing.T) {

	userService := Connect()

	testTable := []struct {
		name     string
		testUser entity.User
		expected error
	}{
		{
			name: "Test Add One",
			testUser: entity.User{
				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: nil,
		},
		{
			name: "Test Add Two",
			testUser: entity.User{
				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Fast",
				Lastname:  "Ganzales",
				Email:     "fast.ganzales.2003@yandex.ru",
				Age:       100,
			},
			expected: nil,
		},
		{
			name: "Test Add Three",
			testUser: entity.User{
				ID:        uuid.MustParse("18eaabba-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "George",
				Lastname:  "Floyd",
				Email:     "george.floyd.2003@breathe.en",
				Age:       100,
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := service.UserRepository.CreateUser(userService, testCase.testUser)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userService := Connect()
	testTable := []struct {
		name     string
		testUser entity.User
		expected error
	}{
		{
			name: "Test Update One",
			testUser: entity.User{
				ID:        uuid.MustParse("18e90062-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Denis",
				Lastname:  "Leonov",
				Email:     "denis.leonov.2003@yandex.ru",
				Age:       19,
			},
			expected: nil,
		},
		{
			name: "Test Update Two",
			testUser: entity.User{
				ID:        uuid.MustParse("18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Andrey",
				Lastname:  "Ivanov",
				Email:     "andrey.ivanov.2003@yandex.ru",
				Age:       19,
			},
			expected: nil,
		},
		{
			name: "Test Update Three",
			testUser: entity.User{
				ID:        uuid.MustParse("18eaabba-54b0-11ed-86a7-e8d8d1f76e0b"),
				Firstname: "Vladimir",
				Lastname:  "Putin",
				Email:     "vladimir.putin.2022@russia.ru",
				Age:       60,
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := service.UserRepository.UpdateUser(userService, testCase.testUser.ID.String(), testCase.testUser)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	userService := Connect()
	testTable := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "Test Delete One",
			id:       "18e90062-54b0-11ed-86a7-e8d8d1f76e0b",
			expected: nil,
		},
		{
			name:     "Test Delete Two",
			id:       "18ea73dc-54b0-11ed-86a7-e8d8d1f76e0b",
			expected: nil,
		},
		{
			name:     "Test Delete Three",
			id:       "18eaabba-54b0-11ed-86a7-e8d8d1f76e0b",
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := service.UserRepository.DeleteUser(userService, testCase.id)

			if actual != testCase.expected {
				t.Errorf("not working. Expect %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	testTable := []struct {
		name     string
		testUser entity.User
		expected error
	}{
		{
			name: "Validation Test One",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: errors.New(""),
		},
		{
			name: "Validation Test Two",
			testUser: entity.User{
				Firstname: "",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: errors.New("и много ты знаешь людей с именем длиной в 1 букву или в 11"),
		},
		{
			name: "Validation Test Three",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "K",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: errors.New("и много ты знаешь людей с фамилией длиной в 1 букву или в 11"),
		},
		{
			name: "Validation Test Four",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020mail.ru",
				Age:       100,
			},
			expected: errors.New("брат, почта стремная, переделай"),
		},
		{
			name: "Validation Test Five",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "badpochta",
				Age:       100,
			},
			expected: errors.New("брат, почта стремная, переделай"),
		},
		{
			name: "Validation Test Six",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "",
				Age:       100,
			},
			expected: errors.New("брат, почта стремная, переделай"),
		},
		{
			name: "Validation Test Seven",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       0,
			},
			expected: errors.New("денис, 0 лет, пошлый. А если серьзно, вводи настоящий возраст, деанона не будет"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := checks.Validation(testCase.testUser)

			if errors.Is(actual, testCase.expected) {
				t.Errorf("not working. Expect %s, got %s", testCase.expected.Error(), actual.Error())
			}
		})
	}
}
