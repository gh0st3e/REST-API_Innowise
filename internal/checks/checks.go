package checks

import (
	"InnowisePreTraineeTask/internal/entity"
	"errors"
	"regexp"
)

func Validation(user entity.User) error {
	if len(user.Firstname) < 2 || len(user.Firstname) > 10 {
		return errors.New("Firstname should be more then 1 symbol but less then 11")
	}
	if len(user.Lastname) < 2 || len(user.Lastname) > 10 {
		return errors.New("Lastname should be more then 1 symbol but less then 11")
	}
	if !validateEmail(user.Email) {
		return errors.New("your email incorrect")
	}
	if user.Age < 1 {
		return errors.New("your should be older tneh 0 year")
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
