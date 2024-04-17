package budget_handler_test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	budget_handler "tracker/pkg/handler/budget"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BudgetnHandler_GetTransactions(t *testing.T) {
	var transactions []model.Transaction
	transaction_1 := mocks.GenerateTransaction(1, 1)
	transaction_2 := mocks.GenerateTransaction(1, 2)

	transactions = append(transactions, transaction_1)
	transactions = append(transactions, transaction_2)
	budget := mocks.GenerateBudget(1)

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		expectedSqlErr error
		budgetId       string
	}{
		{
			name:           "returns transactions",
			expected:       `[{ "amount":6.9, "budgetId":1, "created":"2021-12-12T09:10:00Z", "description":"mock desc", "id":1, "userId":0, "uuid":"mock uuid" }, { "amount":6.9, "budgetId":1, "created":"2021-12-12T09:10:00Z", "description":"mock desc", "id":2, "userId":0, "uuid":"mock uuid" }]`,
			expectedStatus: 200,
			expectedSqlErr: nil,
			budgetId:       "1",
		},
		{
			name:           "returns 404 if transaction not found",
			expected:       `{"msg":"Budget not found"}`,
			expectedStatus: 404,
			expectedSqlErr: sql.ErrNoRows,
			budgetId:       "9999",
		},
		{
			name:           "returns 422 if budgetId not valid",
			expected:       `{"msg":"Invalid Budget id"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			budgetId:       "test",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			columnsBudget := []string{
				"ID",
				"Uuid",
				"Created",
				"Description",
				"Title",
				"UserID",
			}
			columnsTransactions := []string{
				"ID",
				"Uuid",
				"Description",
				"Amount",
				"Created",
				"BudgetID",
				"UserID",
			}
			expectedTransactionsRows := sqlmock.NewRows(columnsTransactions)

			for _, transaction := range transactions {
				expectedTransactionsRows.AddRow(
					transaction.ID,
					transaction.Uuid,
					transaction.Description,
					transaction.Amount,
					transaction.Created,
					transaction.BudgetID,
					transaction.UserID,
				)
			}

			expectedBudgetRows := sqlmock.NewRows(columnsBudget).
				AddRow(
					budget.ID,
					budget.Uuid,
					budget.Created,
					budget.Description,
					budget.Title,
					budget.UserID,
				)

			id, _ := strconv.Atoi(test.budgetId)
			mock.
				ExpectQuery("SELECT").
				WithArgs(id).
				WillReturnRows(expectedBudgetRows).
				WillReturnError(test.expectedSqlErr)

			mock.
				ExpectQuery("SELECT").
				WithArgs(id).
				WillReturnRows(expectedTransactionsRows)

			url := fmt.Sprintf("/transactions/%s", test.budgetId)
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			q := req.URL.Query()
			q.Add("budgetId", test.budgetId)
			req.URL.RawQuery = q.Encode()

			transactionRepo := repository.NewTransactionRepository(db)

			budgetRepo := repository.NewBudgetRepository(db)
			budgetHandler := budget_handler.NewBudgetHandler(budgetRepo, transactionRepo)
			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(budgetHandler.GetTransactions))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
