package repository_test

import (
	"testing"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetRepo_CreateBudget(t *testing.T) {
	testCases := []struct {
		name              string
		expectedErr       error
		expectedSqlErr    error
		budgetId          int
		budgetTitle       string
		budgetDesctiption string
	}{
		{
			name:              "create budget",
			expectedErr:       nil,
			expectedSqlErr:    nil,
			budgetId:          2,
			budgetTitle:       "Test Title",
			budgetDesctiption: "Test Desc",
		},
		{
			name:              "create budget without desc",
			expectedErr:       nil,
			expectedSqlErr:    nil,
			budgetId:          5,
			budgetTitle:       "Test Title",
			budgetDesctiption: "",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer db.Close()

			mock.
				ExpectExec("INSERT INTO budget (Title, Description) VALUES (?, ?)").
				WithArgs(test.budgetTitle, test.budgetDesctiption).
				WillReturnResult(sqlmock.NewResult(int64(test.budgetId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewBudgetRepository(db)

			newBudget := models.Budget{Title: test.budgetTitle, Description: test.budgetDesctiption, ID: test.budgetId}
			newBudgetId, createErr := repo.CreateBudget(newBudget)
			err := mock.ExpectationsWereMet()

			assert.AssertInt(t, int(newBudgetId), test.budgetId)
			assert.AssertError(t, err, test.expectedSqlErr)
			assert.AssertError(t, createErr, test.expectedErr)
		})
	}
}
