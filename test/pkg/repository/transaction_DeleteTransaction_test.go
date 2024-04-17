package repository_test

import (
	"testing"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TransactionRepo_DeleteTransaction(t *testing.T) {
	testCases := []struct {
		name           string
		expectedErr    error
		expectedSqlErr error
		transactionId  int
	}{
		{
			name:           "delete transaction",
			expectedErr:    nil,
			expectedSqlErr: nil,
			transactionId:  1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer db.Close()

			mock.
				ExpectExec("DELETE FROM transaction WHERE ID = ?").
				WithArgs(
					test.transactionId,
				).
				WillReturnResult(sqlmock.NewResult(int64(test.transactionId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewTransactionRepository(db)

			createErr := repo.DeleteTransaction(test.transactionId)
			err := mock.ExpectationsWereMet()

			assert.AssertError(t, err, test.expectedSqlErr)
			assert.AssertError(t, createErr, test.expectedErr)
		})
	}
}
