package handlers_test

import (
	"bytes"
	"database/sql"
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

func Test_TransactionHandler_UpdateTransaction(t *testing.T) {
	transaction := mocks.GenerateTransaction(1, 1)

	testCases := []struct {
		name              string
		expected          string
		expectedStatus    int
		expectedSqlErr    error
		transactionId     string
		transactionAmount float64
		body              string
	}{
		{
			name:           "returns 422 if transactionId not valid",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			expectedSqlErr: nil,
			transactionId:  "test",
			body:           "{}",
		},
		{
			name:           "returns 400 if invalide body",
			expected:       `{"msg": "Invalid request"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			transactionId:  "1",
			body:           "{broken",
		},
		{
			name:           "returns 400 if amoutnt is not valid",
			expected:       `{"msg":"Invalid request"}`,
			expectedStatus: 400,
			expectedSqlErr: nil,
			transactionId:  "1",
			body:           `{"amount": "test"}`,
		},
		{
			name:              "returns 404 if transaction not found",
			expected:          `{"msg":"Transaction not found"}`,
			expectedStatus:    404,
			expectedSqlErr:    sql.ErrNoRows,
			transactionId:     "9999",
			transactionAmount: 12.3,
			body:              `{"amount": 12.3}`,
		},
		{
			name:              "returns msg after update",
			expected:          `{"msg":"Transaction updated"}`,
			expectedStatus:    200,
			expectedSqlErr:    nil,
			transactionId:     "1",
			transactionAmount: 6.9,
			body:              `{"amount": 6.9, "description": ""}`,
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
				ExpectExec("UPDATE").
				WithArgs(test.transactionAmount, "", id).
				WillReturnResult(sqlmock.NewResult(int64(id), 1))

			url := fmt.Sprintf("/transactions/%s", test.transactionId)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(test.body)))

			transactionRepo := repository.NewTransactionRepository(db)
			transactionHandler := handler.NewTransactionHandler(transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(transactionHandler.UpdateTransaction))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
