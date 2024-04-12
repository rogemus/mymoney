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
		Password: "$2a$14$K4TVJJ43ddGnXZ/65J4EyOwGtgTx6UWjDyxmyhPqXWI0qhg0kGXty", // pass
		Email:    "mock@mock.com",
		Username: "Mock Mosinski",
	}
}
