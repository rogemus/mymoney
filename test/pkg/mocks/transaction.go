package mocks_test

import (
	"time"
	"tracker/pkg/model"
)

func GenerateTransaction(budgetID, id int) model.Transaction {
	mock_time := time.Date(2020, 23, 40, 56, 70, 0, 0, time.UTC)

	return model.Transaction{
    ID:          id,
		Uuid:        "mock uuid",
		Description: "mock desc",
		Amount:      6.9,
		Created:     mock_time,
		BudgetID:    budgetID,
	}
}
