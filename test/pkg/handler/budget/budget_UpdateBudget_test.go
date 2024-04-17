package budget_handler_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	budget_handler "tracker/pkg/handler/budget"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetHandler_UpdateBudget(t *testing.T) {
	budget := mocks.GenerateBudget(1)
	testCases := []struct {
		name              string
		expected          string
		expectedStatus    int
		expectedErr       error
		body              string
		budgetId          string
		budgetTitle       string
		budgetDesctiption string
	}{
		{
			name:              "returns msg after updated",
			expected:          `{"msg":"Budget updated"}`,
			expectedStatus:    200,
			body:              `{"title": "Test 1", "description": "Test Desc 1"}`,
			budgetId:          "1",
			budgetTitle:       "Test",
			budgetDesctiption: "Test Desc 1",
		},
		{
			name:              "returns error if empty title",
			expected:          `{"msg":"Invalid request"}`,
			expectedStatus:    400,
			body:              `{"title": "", "description": "Test Desc 2"}`,
			budgetId:          "1",
			budgetTitle:       "",
			budgetDesctiption: "desc",
		},
		{
			name:              "returns error if broken json",
			expected:          `{"msg":"Invalid request"}`,
			expectedStatus:    400,
			body:              `{broken}`,
			budgetId:          "1",
			budgetTitle:       "",
			budgetDesctiption: "desc",
		},
		{
			name:              "returns error if invalid id",
			expected:          `{"msg":"Invalid request"}`,
			expectedStatus:    400,
			body:              `{"title": "Test 4", "description": ""}`,
			budgetId:          "invalid_id",
			budgetTitle:       "Test",
			budgetDesctiption: "",
		},
		{
			name:              "returns error if updating not existing budget",
			expected:          `{"msg":"Budget not found"}`,
			expectedStatus:    404,
			expectedErr:       sql.ErrNoRows,
			body:              `{"title": "Test 4", "description": ""}`,
			budgetId:          "9999",
			budgetTitle:       "Test",
			budgetDesctiption: "",
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
			expectedRows := sqlmock.
				NewRows(columns).
				AddRow(
					&budget.ID,
					&budget.Uuid,
					&budget.Created,
					&budget.Description,
					&budget.Title,
				)

			id, _ := strconv.Atoi(test.budgetId)

			mock.
				ExpectQuery("SELECT").
				WithArgs(id).
				WillReturnError(test.expectedErr).
				WillReturnRows(expectedRows)

			mock.
				ExpectExec("UPDATE").
				WithArgs(test.budgetId, test.budgetTitle, test.budgetDesctiption).
				WillReturnResult(sqlmock.NewResult(int64(id), 1))

			url := fmt.Sprintf("/budget/%s", test.budgetId)
			req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(test.body)))

			budgetRepo := repository.NewBudgetRepository(db)
			transactionRepo := repository.NewTransactionRepository(db)
			budgetHandler := budget_handler.NewBudgetHandler(budgetRepo, transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(budgetHandler.UpdateBudget))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
