package service

import (
	"regexp"
	"tracker/pkg/model"
	errors "tracker/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserValidator struct {
	user model.User
}

func (v *UserValidator) IsEmailValid() (bool, error) {
	email := v.user.Email
	emailRegex, _ := regexp.Compile(`\S+@{1}\S+`)

	if !emailRegex.MatchString(email) {
		return false, errors.UserInvalidEmail
	}

	return true, nil
}

func (v *UserValidator) IsPassValid(hash string) (bool, error) {
	password := v.user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, errors.UserInvalidPassword
	}

	return true, nil
}

func (v *UserValidator) IsUsernameValid() (bool, error) {
	username := v.user.Username

	if username == "" || len(username) > 124 {
		return false, errors.UserInvalidUsername
	}

	return true, nil
}
