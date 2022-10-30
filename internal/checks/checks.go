package checks

import (
	"InnowisePreTraineeTask/internal/entity"
	"errors"
	"regexp"
)

func Validation(user entity.User) error {
	if len(user.Firstname) < 2 || len(user.Firstname) > 10 {
		return errors.New("и много ты знаешь людей с именем длиной в 1 букву или в 11")
	}
	if len(user.Lastname) < 2 || len(user.Lastname) > 10 {
		return errors.New("и много ты знаешь людей с фамилией длиной в 1 букву или в 11")
	}
	if !validateEmail(user.Email) {
		return errors.New("брат, почта стремная, переделай")
	}
	if user.Age < 1 {
		return errors.New("денис, 0 лет, пошлый. А если серьзно, вводи настоящий возраст, деанона не будет")
	}
	return nil
}

func validateEmail(value string) bool {

	var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	if !mailRe.MatchString(value) {
		return false
	}
	return true
}
