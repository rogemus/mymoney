package repository_test

import (
	"testing"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	errors "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func TestGetBudget(t *testing.T) {
	var empty_budget models.Budget
	budget := mocks.GenerateBudget()

	testCases := []struct {
		name          string
		expected      models.Budget
		expectedQuery string
		budgetID      int
		err           error
	}{
		{
			name:          "returns rows for budgetID(1)",
			expected:      budget,
			expectedQuery: "SELECT ID, Uuid, Created, Description, Title FROM budget WHERE ID = ?",
			budgetID:      1,
			err:           nil,
		},
		{
			name:          "returns empty row for budgetID(9999)",
			expected:      empty_budget,
			expectedQuery: "SELECT ID, Uuid, Created, Description, Title FROM budget WHERE ID = ?",
			budgetID:      9999,
			err:           errors.Budget404Err,
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

			if !cmp.Equal(test.expected, empty_budget) {
				bud := test.expected

				expectedRows.AddRow(
					bud.ID,
					bud.Uuid,
					bud.Created,
					bud.Description,
					bud.Title,
				)
			}

			mock.
				ExpectQuery(test.expectedQuery).
				WithArgs(test.budgetID).
				WillReturnError(test.err).
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := repository.NewBudgetRepository(db)
			result, err := repo.GetBudget(test.budgetID)

			assert.AssertStruct[models.Budget](t, result, test.expected)
			assert.AssertError(t, err, test.err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
