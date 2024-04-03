package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
	"tracker/pkg/models"
	assert "tracker/utils/testing"
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
	var transactions []models.Transaction

	transaction := generateTransactions(1)
	transactions = append(transactions, transaction)
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
		transaction.TransactionID,
		transaction.TransactionUuid,
		transaction.Description,
		transaction.Amount,
		transaction.Created,
		transaction.BudgetID,
	)

	t.Run("returns rows for budget", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		budgetID := 1
		mock.
			ExpectQuery(QueryGetTransactions).
			WithArgs(budgetID).
			WillReturnRows(expectedRows)

		if err != nil {
			t.Fatalf("an error has occured: %s", err)
		}

		defer db.Close()

		repo := NewTransactionRepository(db)
		result, err := repo.GetTransactions(budgetID)

		assert.AssertEqualInt(t, len(result), len(transactions))
		assert.AssertSliceOfStructs(t, result, transactions)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("return empty rows", func(t *testing.T) {
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
		// expectedRows.AddRow(
		// 	transaction.TransactionID,
		// 	transaction.TransactionUuid,
		// 	transaction.Description,
		// 	transaction.Amount,
		// 	transaction.Created,
		// 	transaction.BudgetID,
		// )
		budgetID := 99999
		mock.
			ExpectQuery(QueryGetTransactions).
			WithArgs(budgetID).
			WillReturnRows(expectedRows)

		if err != nil {
			t.Fatalf("an error has occured: %s", err)
		}

		defer db.Close()

		repo := NewTransactionRepository(db)
		transactions, err := repo.GetTransactions(budgetID)

		assert.AssertEqualInt(t, len(transactions), 4)
		assert.AssertSliceOfStructs[models.Transaction](t, transactions, make([]models.Transaction, 0))

		if err != nil {
			t.Error(err)
		}
	})
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
