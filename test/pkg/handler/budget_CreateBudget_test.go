package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"tracker/pkg/handler"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetHandler_CreateBudget(t *testing.T) {
	testCases := []struct {
		name              string
		expected          string
		expectedStatus    int
		body              string
		budgetTitle       string
		budgetDesctiption string
	}{
		{
			name:              "returns msg after create",
			expected:          `{"msg":"Budget created"}`,
			expectedStatus:    201,
			body:              `{"title": "Test 1", "description": "Test Desc 1"}`,
			budgetTitle:       "Test",
			budgetDesctiption: "Desc",
		},
		{
			name:              "returns error if broken json",
			expected:          `{"msg":"Invalid request"}`,
			expectedStatus:    400,
			body:              `{broken}`,
			budgetTitle:       "",
			budgetDesctiption: "desc",
		},
		{
			name:              "returns error if empty title",
			expected:          `{"msg":"Invalid request"}`,
			expectedStatus:    400,
			body:              `{"title": "", "description": "Test Desc 2"}`,
			budgetTitle:       "",
			budgetDesctiption: "desc",
		},
		{
			name:              "returns msg after create when empty desc",
			expected:          `{"msg":"Budget created"}`,
			expectedStatus:    201,
			body:              `{"title": "Test 4", "description": ""}`,
			budgetTitle:       "Test",
			budgetDesctiption: "",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()

			mock.
				ExpectExec("INSERT").
				WithArgs(test.budgetTitle, test.budgetDesctiption).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req := httptest.NewRequest(http.MethodPost, "/budget", bytes.NewBuffer([]byte(test.body)))

			budgetRepo := repository.NewBudgetRepository(db)
			transactionRepo := repository.NewTransactionRepository(db)
			budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(budgetHandler.CreateBudget))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
