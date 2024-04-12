package mocks_test

import (
	"time"
	"tracker/pkg/model"
)

func GenerateToken() model.Token {
	return model.Token{
		ID:        1,
		Uuid:      "uuid",
		Created:   time.Date(2020, 23, 40, 56, 70, 0, 0, time.UTC).UTC(),
		Token:     "token.token.token",
		UserEmail: "user@user.com",
	}
}
