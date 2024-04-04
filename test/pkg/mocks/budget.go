package mocks_test

import (
	"fmt"
	"math/rand"
	"time"
	"tracker/pkg/models"
)

func GenerateBudget(id ...int) models.Budget {
	mock_time := time.Date(2020, 23, 40, 56, 70, 0, 0, time.UTC)
	mock_id := rand.Intn(9999)

	if id != nil {
		mock_id = id[0]
	}

	return models.Budget{
		ID:          mock_id,
		Uuid:        fmt.Sprintf("%d", rand.Intn(9999)),
		Created:     mock_time.UTC(),
		Description: fmt.Sprintf("description %d", rand.Intn(9999)),
		Title:       fmt.Sprintf("title %d", rand.Intn(9999)),
	}
}
