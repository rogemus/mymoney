package repository_test

import (
	"database/sql"
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TransactionRepo_GetTransaction(t *testing.T) {
	transaction := mocks.GenerateTransaction(1, 1)
	var empty_transaction model.Transaction

	testCases := []struct {
		name           string
		expected       model.Transaction
		transactionID  int
		expectedSqlErr error
	}{
		{
			name:           "returns row",
			expected:       transaction,
			transactionID:  1,
			expectedSqlErr: nil,
		},
		{
			name:           "returns err if not found",
			expected:       empty_transaction,
			transactionID:  9999,
			expectedSqlErr: sql.ErrNoRows,
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

			expectedRows.AddRow(
				transaction.ID,
				transaction.Uuid,
				transaction.Description,
				transaction.Amount,
				transaction.Created,
				transaction.BudgetID,
				transaction.UserID,
			)

			mock.
				ExpectQuery("SELECT ID, Uuid, Description, Amount, Created, BudgetID, UserID FROM transaction WHERE ID = ?").
				WithArgs(test.transactionID).
				WillReturnRows(expectedRows).
				WillReturnError(test.expectedSqlErr)
			defer db.Close()

			repo := repository.NewTransactionRepository(db)
			result, getErr := repo.GetTransaction(test.transactionID)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, test.expectedSqlErr)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertInt(t, result.ID, test.expected.ID)
			assert.AssertStruct[model.Transaction](t, result, test.expected)
		})
	}
}
