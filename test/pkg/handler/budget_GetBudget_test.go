package handlers_test

import (
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

func Test_BudgetHandler_GetBudget(t *testing.T) {
	budget := mocks.GenerateBudget(1)

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		budgetId       string
		expectedErr    error
	}{
		{
			name:           "returns budget json",
			expected:       `{"budget":{"id":1,"uuid":"mock uuid","created":"2021-12-12T09:10:00Z","description":"mock description","title":"mock title", "userId": 1},"transactions":null}`,
			expectedStatus: 200,
			budgetId:       "1",
			expectedErr:    nil,
		},
		{
			name:           "returns 404 if budget not found",
			expected:       `{"msg":"Budget not found"}`,
			expectedStatus: 404,
			budgetId:       "9999",
			expectedErr:    sql.ErrNoRows,
		},
		{
			name:           "returns 400 if budgetId not valid",
			expected:       `{"msg":"Invalid request"}`,
			expectedStatus: 400,
			budgetId:       "test",
			expectedErr:    nil,
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
			expectedRows := sqlmock.
				NewRows(columns).
				AddRow(
					budget.ID,
					budget.Uuid,
					budget.Created,
					budget.Description,
					budget.Title,
					budget.UserID,
				)

			id, _ := strconv.Atoi(test.budgetId)

			if test.expectedErr != nil {
				mock.
					ExpectQuery("SELECT").
					WithArgs(id).
					WillReturnError(sql.ErrNoRows)
			} else {
				mock.
					ExpectQuery("SELECT").
					WithArgs(id).
					WillReturnRows(expectedRows)
			}

			url := fmt.Sprintf("/budget/%s", test.budgetId)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			budgetRepo := repository.NewBudgetRepository(db)
			transactionRepo := repository.NewTransactionRepository(db)
			budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(mocks.MockProtected(budgetHandler.GetBudget))
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
