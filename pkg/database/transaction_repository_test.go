package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
	"tracker/pkg/models"
)

func generateTransactions(budgetID int) models.Transaction {
	return models.Transaction{
		TransactionUuid: fmt.Sprintf("%d", rand.Intn(9999)),
		TransactionID:   rand.Intn(9999),
		Description:     fmt.Sprintf("description %d", rand.Intn(9999)),
		Amount:          rand.Float32(),
		Created:         time.Now(),
		BudgetID:        budgetID,
	}
}

func TestGetTransactions(t *testing.T) {
	var tts []models.Transaction
	tt := generateTransactions(1)
	tts = append(tts, tt)

	t.Run("returns rows for budget", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		columns := []string{
			"TransactionID",
			"TransactionUuid",
			"Description",
			"Amount",
			"Created",
			"BudgetID",
		}
		expectedRows := sqlmock.NewRows(columns)
		expectedRows.AddRow(
			tt.TransactionID,
			tt.TransactionUuid,
			tt.Description,
			tt.Amount,
			tt.Created,
			tt.BudgetID,
		)
		mock.
			ExpectQuery(QueryGetTransactions).
			WithArgs(1).
			WillReturnRows(expectedRows)

		if err != nil {
			t.Fatalf("an error has occured: %s", err)
		}

		defer db.Close()

		repo := NewTransactionRepository(db)
		tts, err = repo.GetTransactions(1)

		assertEqualInt(t, len(tts), 1)

		if err != nil {
			t.Error(err)
		}
	})
}

func assertEqualInt(t testing.TB, got int, want int) {
	t.Helper()

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

// func TestGetTransactions(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	tr := NewTransactionRepository(db)
// 	defer db.Close()
//
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	t.Run("do not return rows", func(t *testing.T) {
// 		// rows := sqlmock.NewRows([]string{"TransactionID", "BudgetID"}).AddRow(1, 3)
// 		budgetId := 5
// 		query := `SELECT \* FROM transaction WHERE BudgetID = \?;`
//
// 		tr.GetTransactions(budgetId)
//
// 		mock.ExpectBegin()
// 		mock.
// 			ExpectQuery(query).
// 			WithArgs("test")
//
// 		// 	WithArgs(1).
// 		// 	WillReturnRows(rows)
//
// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("ther were unfulfiled expectations: %s", err)
// 		}
// 	})
// }
