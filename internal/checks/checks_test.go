package checks

import (
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
			expected: errors.New("Firstname should be more then 1 symbol but less then 11"),
		},
		{
			name: "Validation Test Three",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "K",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       100,
			},
			expected: errors.New("Lastname should be more then 1 symbol but less then 11"),
		},
		{
			name: "Validation Test Four",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020mail.ru",
				Age:       100,
			},
			expected: errors.New("your email incorrect"),
		},
		{
			name: "Validation Test Five",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "badpochta",
				Age:       100,
			},
			expected: errors.New("your email incorrect"),
		},
		{
			name: "Validation Test Six",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "",
				Age:       100,
			},
			expected: errors.New("your email incorrect"),
		},
		{
			name: "Validation Test Seven",
			testUser: entity.User{
				Firstname: "Viktor",
				Lastname:  "Korneplod",
				Email:     "viktor.korneplod.2020@mail.ru",
				Age:       0,
			},
			expected: errors.New("your should be older tneh 0 year"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Validation(testCase.testUser)

			if errors.Is(actual, testCase.expected) {
				t.Errorf("not working. Expect %s, got %s", testCase.expected.Error(), actual.Error())
			}
		})
	}
}
