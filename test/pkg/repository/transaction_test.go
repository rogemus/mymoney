package repository_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
)

func generateTransactions(budgetID int) models.Transaction {
	return models.Transaction{
		Uuid:        fmt.Sprintf("%d", rand.Intn(9999)),
		ID:          rand.Intn(9999),
		Description: fmt.Sprintf("description %d", rand.Intn(9999)),
		Amount:      rand.Float32(),
		Created:     time.Now(),
		BudgetID:    budgetID,
	}
}

break tests

func TestGetTransactions(t *testing.T) {
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
			expectedQuery: "SELECT * FROM transaction WHERE BudgetID = ?",
			budgetID:      1,
		},
		{
			name:          "returns empty row for budgetID(9999)",
			expected:      make([]models.Transaction, 0),
			expectedQuery: "SELECT * FROM transaction WHERE BudgetID = ?",
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
					transaction.ID,
					transaction.Uuid,
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

			repo := repository.NewTransactionRepository(db)
			result, err := repo.GetTransactionsForBudget(test.budgetID)

			assert.AssertEqualInt(t, len(result), len(test.expected))
			assert.AssertSliceOfStructs(t, result, test.expected)

			if err != nil {
				t.Error(err)
			}
		})
	}
}
