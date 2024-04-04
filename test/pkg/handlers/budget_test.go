package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tracker/pkg/handler"
	"tracker/pkg/repository"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetBudget(t *testing.T) {
	budget := mocks.GenerateBudget(1)
	db, mock, _ := sqlmock.New()
	//
	columns := []string{
		"ID",
		"Uuid",
		"Created",
		"Description",
		"Title",
	}
	expectedRows := sqlmock.NewRows(columns).
		AddRow(
			budget.ID,
			budget.Uuid,
			budget.Created,
			budget.Description,
			budget.Title,
		)

	mock.ExpectQuery("SELECT").WillReturnRows(expectedRows)
	req, _ := http.NewRequest("GET", "/budget/1", nil)

	budgetRepo := repository.NewBudgetRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

	rr := httptest.NewRecorder()
	hr := http.HandlerFunc(budgetHandler.GetBudget)

	hr.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
