package repository_test

import (
	"testing"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetRepo_DeleteBudget(t *testing.T) {
	testCases := []struct {
		name           string
		budgetID       int
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "delete budget",
			budgetID:       1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "delete budget if not exist",
			budgetID:       9999,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			defer db.Close()

			mock.
				ExpectExec("DELETE FROM budget WHERE ID = ?").
				WithArgs(test.budgetID).
				WillReturnResult(sqlmock.NewResult(int64(test.budgetID), 1)).
				WillReturnError(test.expectedSqlErr)
			defer db.Close()

			repo := repository.NewBudgetRepository(db)
			delErr := repo.DeleteBudget(test.budgetID)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, delErr, nil)
			assert.AssertError(t, sqlErr, nil)
		})
	}
}
