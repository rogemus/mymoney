package repository_test

import (
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TransactionRepo_GetTransactionsForBudget(t *testing.T) {
	var transactions []model.Transaction

	transaction_1 := mocks.GenerateTransaction(1, 1)
	transactions = append(transactions, transaction_1)

	testCases := []struct {
		name           string
		expected       []model.Transaction
		budgetID       int
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "returns rows for budgetID(1)",
			expected:       transactions,
			budgetID:       1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "returns empty row for budgetID(9999)",
			expected:       make([]model.Transaction, 0),
			budgetID:       9999,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"ID",
				"Uuid",
				"Description",
				"Amount",
				"Created",
				"BudgetID",
				"UserID",
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
					transaction.UserID,
				)
			}

			mock.
				ExpectQuery("SELECT Description, Amount, UserID FROM transaction WHERE BudgetID = ?").
				WithArgs(test.budgetID).
				WillReturnRows(expectedRows).
				WillReturnError(test.expectedSqlErr)

			defer db.Close()

			repo := repository.NewTransactionRepository(db)
			result, getErr := repo.GetTransactionsForBudget(test.budgetID)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertInt(t, len(result), len(test.expected))
			assert.AssertSliceOfStructs(t, result, test.expected)
			assert.AssertError(t, getErr, test.expectedErr)
			assert.AssertError(t, sqlErr, nil)
		})
	}
}
