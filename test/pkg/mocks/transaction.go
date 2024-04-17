package mocks_test

import (
	"fmt"
	"math/rand"
	"time"
	"tracker/pkg/model"
)

func GenerateTransaction(budgetID, id int) model.Transaction {
	mock_time := time.Date(2020, 23, 40, 56, 70, 0, 0, time.UTC)

	return model.Transaction{
    ID:          id,
		Uuid:        fmt.Sprintf("%d", rand.Intn(9999)),
		Description: fmt.Sprintf("description %d", rand.Intn(9999)),
		Amount:      rand.Float64(),
		Created:     mock_time,
		BudgetID:    budgetID,
	}
}
