package tests

import (
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/repository"
	"InnowisePreTraineeTask/internal/service"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCreateUser(t *testing.T) {
	log := logrus.New()
	db := database.Connect()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(log, userRepository)

	testTable := []struct {
		testUser entity.User
		expected error
	}{
		{
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: nil,
		},
		{
			testUser: entity.User{
				Firstname: "Fast",
				Lastname:  "Ganzales",
				Email:     "fast.ganzales.2003@yandex.ru",
				Age:       100,
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		result := service.UserRepository.CreateUser(userService, testCase.testUser)

		if result != testCase.expected {
			t.Errorf("not working. Expect %v, got %v", testCase.expected, result)
		}
	}
}
