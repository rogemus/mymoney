package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"tracker/pkg/handler"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_TransactionHandler_DeleteTransaction(t *testing.T) {
	transaction := mocks.GenerateTransaction(1, 1)

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		expectedSqlErr error
		transactionId  string
	}{
		{
			name:           "returns msg after delete",
			expected:       `{"msg":"Transaction deleted"}`,
			expectedStatus: 204,
			expectedSqlErr: nil,
			transactionId:  "1",
		},
		{
			name:           "returns 404 if transaction not found",
			expected:       `{"msg":"Transaction not found"}`,
			expectedStatus: 404,
			expectedSqlErr: nil,
			transactionId:  "9999",
		},
		{
			name:           "returns 422 if transactionId not valid",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			expectedSqlErr: nil,
			transactionId:  "test",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			columns := []string{
				"ID",
				"Uuid",
				"Description",
				"Amount",
				"Created",
				"BudgetID",
				"UserID",
			}
			expectedRows := sqlmock.NewRows(columns)

			expectedRows.AddRow(
				transaction.ID,
				transaction.Uuid,
				transaction.Description,
				transaction.Amount,
				transaction.Created,
				transaction.BudgetID,
				transaction.UserID,
			)

			id, _ := strconv.Atoi(test.transactionId)

			mock.
				ExpectQuery("SELECT").
				WithArgs(id).
				WillReturnRows(expectedRows).
				WillReturnError(test.expectedSqlErr)

			mock.
				ExpectExec("DELETE").
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(int64(id), 1))

			url := fmt.Sprintf("/transactions/%s", test.transactionId)
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			transactionRepo := repository.NewTransactionRepository(db)
			transactionHandler := handler.NewTransactionHandler(transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(transactionHandler.DeleteTransaction))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
