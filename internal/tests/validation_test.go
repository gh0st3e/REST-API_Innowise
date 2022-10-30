package tests

import (
	"InnowisePreTraineeTask/internal/checks"
	"InnowisePreTraineeTask/internal/entity"
	"errors"
	"testing"
)

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
