package repository_test

import (
	"testing"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepoGetTransactions(t *testing.T) {
	var transactions []models.Transaction

	transaction_1 := mocks.GenerateTransaction(1)
	transactions = append(transactions, transaction_1)

	testCases := []struct {
		name     string
		expected []models.Transaction
		budgetID int
	}{
		{
			name:     "returns rows for budgetID(1)",
			expected: transactions,
			budgetID: 1,
		},
		{
			name:     "returns empty row for budgetID(9999)",
			expected: make([]models.Transaction, 0),
			budgetID: 9999,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"ID",
				"Uuid",
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
				ExpectQuery("SELECT * FROM transaction WHERE BudgetID = ?").
				WithArgs(test.budgetID).
				WillReturnRows(expectedRows)

			if err != nil {
				t.Fatalf("an error has occured: %s", err)
			}

			defer db.Close()

			repo := repository.NewTransactionRepository(db)
			result, err := repo.GetTransactionsForBudget(test.budgetID)

			assert.AssertInt(t, len(result), len(test.expected))
			assert.AssertSliceOfStructs(t, result, test.expected)

			if err != nil {
				t.Error(err)
			}
		})
	}
}
