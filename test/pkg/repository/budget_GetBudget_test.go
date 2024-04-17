package repository_test

import (
	"database/sql"
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func Test_BudgetRepo_GetBudget(t *testing.T) {
	var empty_budget model.Budget
	budget := mocks.GenerateBudget()

	testCases := []struct {
		name           string
		expected       model.Budget
		budgetID       int
		expectedSqlErr error
	}{
		{
			name:           "returns rows for budgetID(1)",
			expected:       budget,
			budgetID:       1,
			expectedSqlErr: nil,
		},
		{
			name:           "returns empty row for budgetID(9999)",
			expected:       empty_budget,
			budgetID:       9999,
			expectedSqlErr: sql.ErrNoRows,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"ID",
				"Uuid",
				"Created",
				"Description",
				"Title",
				"UserID",
			}
			expectedRows := sqlmock.NewRows(columns)

			if !cmp.Equal(test.expected, empty_budget) {
				bud := test.expected

				expectedRows.AddRow(
					bud.ID,
					bud.Uuid,
					bud.Created,
					bud.Description,
					bud.Title,
					bud.UserID,
				)
			}

			mock.
				ExpectQuery("SELECT ID, Uuid, Created, Description, Title, UserID FROM budget WHERE ID = ?").
				WithArgs(test.budgetID).
				WillReturnError(test.expectedSqlErr).
				WillReturnRows(expectedRows)
			defer db.Close()

			repo := repository.NewBudgetRepository(db)
			result, getErr := repo.GetBudget(test.budgetID)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, test.expectedSqlErr)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertStruct[model.Budget](t, result, test.expected)
		})
	}
}
