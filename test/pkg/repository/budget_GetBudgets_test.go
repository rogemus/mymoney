package repository_test

import (
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetRepo_GetBudgets(t *testing.T) {
	var budgets []model.Budget
	var empty_budgets []model.Budget
	budget_1 := mocks.GenerateBudget(1)
	budget_2 := mocks.GenerateBudget(2)

	budgets = append(budgets, budget_1)
	budgets = append(budgets, budget_2)

	testCases := []struct {
		name           string
		expected       []model.Budget
		expectedSqlErr error
	}{
		{
			name:           "returns rows for budgets",
			expected:       budgets,
			expectedSqlErr: nil,
		},
		{
			name:           "returns empty row for budgets",
			expected:       empty_budgets,
			expectedSqlErr: nil,
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

			for _, budget := range test.expected {
				expectedRows.AddRow(
					budget.ID,
					budget.Uuid,
					budget.Created,
					budget.Description,
					budget.Title,
					budget.UserID,
				)
			}

			mock.
				ExpectQuery("SELECT ID, Uuid, Created, Description, Title, UserID FROM budget").
				WithoutArgs().
				WillReturnRows(expectedRows)
			defer db.Close()

			repo := repository.NewBudgetRepository(db)
			result, getErr := repo.GetBudgets()
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, nil)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertSliceOfStructs[model.Budget](t, result, test.expected)
		})
	}
}
