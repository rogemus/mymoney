package service

import (
	"regexp"
	"tracker/pkg/model"
	errors "tracker/pkg/utils"
)

func ValidateUser(user model.User) error {
	emailRegex, _ := regexp.Compile(`\S+@{1}\S+`)

	if user.Username == "" || len(user.Username) > 124 {
		return errors.UserInvalidUsername
	}

	if !emailRegex.MatchString(user.Email) {
		return errors.UserInvalidEmail
	}

	// TODO: validate how strong is pass
	if user.Password == "" {
		return errors.UserInvalidPassword
	}

	return errors.Generic400Err
}
