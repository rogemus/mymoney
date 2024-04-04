package repository

import (
	"database/sql"
	"tracker/pkg/models"
)

type TransactionRepository interface {
	GetTransactionsForBudget(budgetId int) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetTransactionsForBudget(budgetId int) ([]models.Transaction, error) {
	query := "SELECT * FROM transaction WHERE BudgetID = ?"
	rows, err := r.db.Query(query, budgetId)

  if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.Uuid,
			&transaction.Description,
			&transaction.Amount,
			&transaction.Created,
			&transaction.BudgetID,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
