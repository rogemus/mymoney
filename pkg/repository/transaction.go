package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type TransactionRepository interface {
	GetTransactionsForBudget(budgetId int) ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) (int64, error)
	// UpdateTransaction(transaction model.Transaction) error
	// DeleteTransaction(id int) (int64, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction model.Transaction) (int64, error) {
	query := "INSERT INTO transaction (Description, Amount, BudgetID, UserID) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(
		query,
		transaction.Description,
		transaction.Amount,
		transaction.BudgetID,
		transaction.UserID,
	)

	if err != nil {
		return -1, errs.Generic400Err
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		return -1, errs.Generic400Err
	}

	return insertedId, nil
}

// func UpdateTransaction(budgetId, userId int, transaction model.Transaction) (int64, error) {
//
// }
//
// func DeleteTransaction(id int) (int64, error) {
//
// }

func (r *transactionRepository) GetTransactionsForBudget(budgetId int) ([]model.Transaction, error) {
	query := "SELECT Description, Amount, UserID FROM transaction WHERE BudgetID = ?"
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
			&transaction.UserID,
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
