package database

import (
	"database/sql"
	"fmt"
	"tracker/pkg/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (tr *TransactionRepository) GetTransactions(budgetId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `
    SELECT 
      TransactionId, TransactionUuid, Description, Amount, Created, BudgetID
    FROM
      transaction
    WHERE
      BudgetID = ?;
  `
	rows, err := tr.db.Query(query, budgetId)

	if err != nil {
		return nil, fmt.Errorf("GetTransactions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.TransactionID, &t.TransactionUuid, &t.Description, &t.Amount, &t.Created, &t.BudgetID); err != nil {
			return nil, fmt.Errorf("GetTransactions: %v", err)
		}

		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTransactions: %v", err)
	}

	return transactions, nil
}
