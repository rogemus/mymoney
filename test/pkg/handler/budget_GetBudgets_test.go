package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tracker/pkg/handler"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetHandler_GetBudgets(t *testing.T) {
	var budgets []models.Budget
	var empty_budgets []models.Budget
	budget_1 := mocks.GenerateBudget(1)
	budget_2 := mocks.GenerateBudget(2)

	budgets = append(budgets, budget_1)
	budgets = append(budgets, budget_2)

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		expectedErr    error
		expectedData   []models.Budget
	}{
		{
			name:           "returns budgets json",
			expected:       `[{"id":1,"uuid":"mock uuid","created":"2021-12-12T09:10:00Z","description":"mock description","title":"mock title"},{"id":2,"uuid":"mock uuid","created":"2021-12-12T09:10:00Z","description":"mock description","title":"mock title"}]`,
			expectedStatus: 200,
			expectedErr:    nil,
			expectedData:   budgets,
		},
		{
			name:           "returns empty array in json",
			expected:       `[]`,
			expectedStatus: 200,
			expectedErr:    nil,
			expectedData:   empty_budgets,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			columns := []string{
				"ID",
				"Uuid",
				"Created",
				"Description",
				"Title",
			}

			expectedRows := sqlmock.NewRows(columns)

			for _, budget := range test.expectedData {
				expectedRows.AddRow(
					&budget.ID,
					&budget.Uuid,
					&budget.Created,
					&budget.Description,
					&budget.Title,
				)
			}

			mock.
				ExpectQuery("SELECT").
				WillReturnRows(expectedRows)

			req := httptest.NewRequest(http.MethodGet, "/budgets", nil)

			budgetRepo := repository.NewBudgetRepository(db)
			transactionRepo := repository.NewTransactionRepository(db)
			budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(budgetHandler.GetBudgets)
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
