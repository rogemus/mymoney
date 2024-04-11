package service

import (
	"regexp"
	"tracker/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

type UserValidator struct {
	User model.User
}

func (v *UserValidator) IsEmailValid() bool {
	email := v.User.Email
	emailRegex, _ := regexp.Compile(`\S+@{1}\S+`)
	return emailRegex.MatchString(email)
}

func (v *UserValidator) IsPassValid(hash string) bool {
	password := v.User.Password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

func (v *UserValidator) IsUsernameValid() bool {
	username := v.User.Username
	return username != "" || len(username) < 124
}
