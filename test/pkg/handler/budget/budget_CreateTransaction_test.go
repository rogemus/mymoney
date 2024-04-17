package budget_handler_test

import (
	"bytes"
	"database/sql"
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

func Test_BudgetHandler_CreateTransaction(t *testing.T) {
	budget := mocks.GenerateBudget(1)

	testCases := []struct {
		name              string
		expected          string
		expectedStatus    int
		expectedSqlErr    error
		budgetId          string
		transactionAmount float64
		body              string
	}{
		{
			name:           "returns 400 if invalide body",
			expected:       `{"msg": "Invalid request"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			budgetId:       "1",
			body:           "{broken",
		},
		{
			name:           "returns 400 if amount is not valid",
			expected:       `{"msg":"Invalid request"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			budgetId:       "1",
			body:           `{"amount": "test"}`,
		},
		{
			name:           "returns 400 if budgetID",
			expected:       `{"msg":"Invalid Budget id"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			budgetId:       "test",
			body:           `{"amount": 1.2}`,
		},
		{
			name:           "returns 404 if budget not found",
			expected:       `{"msg":"Budget not found"}`,
			expectedStatus: 404,
			expectedSqlErr: sql.ErrNoRows,
			budgetId:       "1",
			body:           `{"budgetId": 9999, "amount": 1.2}`,
		},
		{
			name:              "returns msg after created",
			expected:          `{"msg":"Transaction created"}`,
			expectedStatus:    201,
			expectedSqlErr:    nil,
			transactionAmount: 6.9,
			budgetId:          "1",
			body:              `{"amount": 6.9, "description": ""}`,
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
				"UserID",
			}
			id, _ := strconv.Atoi(test.budgetId)
			expectedRows := sqlmock.NewRows(columns)
			expectedRows.AddRow(
				budget.ID,
				budget.Uuid,
				budget.Created,
				budget.Description,
				budget.Title,
				budget.UserID,
			)

			mock.
				ExpectQuery("SELECT").
				WithArgs(id).
				WillReturnRows(expectedRows).
				WillReturnError(test.expectedSqlErr)

			mock.
				ExpectExec("INSERT").
				WithArgs("", test.transactionAmount, id, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer([]byte(test.body)))
			q := req.URL.Query()
			q.Add("budgetId", test.budgetId)
			req.URL.RawQuery = q.Encode()

			transactionRepo := repository.NewTransactionRepository(db)
			budgetRepo := repository.NewBudgetRepository(db)
			budgetHandler := budget_handler.NewBudgetHandler(budgetRepo, transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(budgetHandler.CreateTransation))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
