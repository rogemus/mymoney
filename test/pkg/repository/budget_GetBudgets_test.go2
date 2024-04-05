package repository_test

import (
	"testing"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetRepo_GetBudgets(t *testing.T) {
	var budgets []models.Budget
	var empty_budgets []models.Budget
	budget_1 := mocks.GenerateBudget(1)
	budget_2 := mocks.GenerateBudget(2)

	budgets = append(budgets, budget_1)
	budgets = append(budgets, budget_2)

	testCases := []struct {
		name           string
		expected       []models.Budget
		expectedQuery  string
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "returns rows for budgets",
			expected:       budgets,
			expectedQuery:  "SELECT ID, Uuid, Created, Description, Title FROM budget",
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "returns empty row for budgets",
			expected:       empty_budgets,
			expectedQuery:  "SELECT ID, Uuid, Created, Description, Title FROM budget",
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"ID",
				"Uuid",
				"Created",
				"Description",
				"Title",
			}
			expectedRows := sqlmock.NewRows(columns)

			for _, budget := range test.expected {
				expectedRows.AddRow(
					budget.ID,
					budget.Uuid,
					budget.Created,
					budget.Description,
					budget.Title,
				)
			}

			mock.
				ExpectQuery(test.expectedQuery).
        WithoutArgs().
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := repository.NewBudgetRepository(db)
			result, err := repo.GetBudgets()

			assert.AssertSliceOfStructs[models.Budget](t, result, test.expected)
			assert.AssertError(t, err, test.expectedErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
