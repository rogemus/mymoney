package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
	"tracker/pkg/models"
	assert "tracker/utils/testing"
)

func generateTransactions(budgetID int) models.Transaction {
	return models.Transaction{
		TransactionUuid: fmt.Sprintf("%d", rand.Intn(9999)),
		TransactionID:   rand.Intn(9999),
		Description:     fmt.Sprintf("description %d", rand.Intn(9999)),
		Amount:          rand.Float32(),
		Created:         time.Now(),
		BudgetID:        budgetID,
	}
}

func TestGetTransactions_2(t *testing.T) {

	var transactions []models.Transaction

	transaction_1 := generateTransactions(1)
	transactions = append(transactions, transaction_1)

	testCases := []struct {
		name          string
		expected      []models.Transaction
		expectedQuery string
		budgetID      int
	}{
		{
			name:          "returns rows for budgetID(1)",
			expected:      transactions,
			expectedQuery: QueryGetTransactions,
			budgetID:      1,
		},
		{
			name:          "returns empty row for budgetID(9999)",
			expected:      make([]models.Transaction, 0),
      expectedQuery: QueryGetTransactions,
			budgetID:      9999,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"TransactionID",
				"TransactionUuid",
				"Description",
				"Amount",
				"Created",
				"BudgetID",
			}
			expectedRows := sqlmock.NewRows(columns)

			for _, transaction := range test.expected {
				expectedRows.AddRow(
					transaction.TransactionID,
					transaction.TransactionUuid,
					transaction.Description,
					transaction.Amount,
					transaction.Created,
					transaction.BudgetID,
				)
			}

			mock.
				ExpectQuery(test.expectedQuery).
				WithArgs(test.budgetID).
				WillReturnRows(expectedRows)

			if err != nil {
				t.Fatalf("an error has occured: %s", err)
			}

			defer db.Close()

			repo := NewTransactionRepository(db)
			result, err := repo.GetTransactions(test.budgetID)

			assert.AssertEqualInt(t, len(result), len(test.expected))
			assert.AssertSliceOfStructs(t, result, test.expected)

			if err != nil {
				t.Error(err)
			}
		})
	}
}
