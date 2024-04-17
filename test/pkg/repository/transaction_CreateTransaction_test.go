package repository_test

import (
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TransactionRepo_CreateTransaction(t *testing.T) {
	transaction := mocks.GenerateTransaction(1, 1)
	var empty_transaction model.Transaction
	testCases := []struct {
		name           string
		expectedErr    error
		expectedSqlErr error
		transaction    model.Transaction
	}{
		{
			name:           "create transaction",
			expectedErr:    nil,
			expectedSqlErr: nil,
			transaction:    transaction,
		},
		{
			name:           "create empty transaction",
			expectedErr:    nil,
			expectedSqlErr: nil,
			transaction:    empty_transaction,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer db.Close()

			mock.
				ExpectExec("INSERT INTO transaction (Description, Amount, BudgetID, UserID) VALUES (?, ?, ?, ?)").
				WithArgs(
					test.transaction.Description,
					test.transaction.Amount,
					test.transaction.BudgetID,
					test.transaction.UserID,
				).
				WillReturnResult(sqlmock.NewResult(int64(test.transaction.ID), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewTransactionRepository(db)

			newTransactionId, createErr := repo.CreateTransaction(test.transaction)
			err := mock.ExpectationsWereMet()

			assert.AssertInt(t, int(newTransactionId), test.transaction.ID)
			assert.AssertError(t, err, test.expectedSqlErr)
			assert.AssertError(t, createErr, test.expectedErr)
		})
	}
}
