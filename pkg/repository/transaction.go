package repository

import (
	"database/sql"
	"tracker/pkg/model"
)

type TransactionRepository interface {
	GetTransactionsForBudget(budgetId int) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetTransactionsForBudget(budgetId int) ([]model.Transaction, error) {
	query := "SELECT * FROM transaction WHERE BudgetID = ?"
	rows, err := r.db.Query(query, budgetId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := make([]model.Transaction, 0)

	for rows.Next() {
		var transaction model.Transaction

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
