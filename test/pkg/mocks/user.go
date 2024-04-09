package mocks_test

import (
	"math/rand"
	"time"
	"tracker/pkg/model"
)

func GenerateUser(id ...int) model.User {
	mock_id := rand.Intn(9999)

	if id != nil {
		mock_id = id[0]
	}

	return model.User{
		ID:       mock_id,
		Uuid:     "mock uuid",
		Created:  time.Date(2020, 23, 40, 56, 70, 0, 0, time.UTC).UTC(),
		Password: "pass",
		Email:    "mock@mock.com",
		Username: "Mock Mosinski",
	}
}
