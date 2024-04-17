package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type TransactionRepository interface {
	GetTransactionsForBudget(budgetId int) ([]model.Transaction, error)
	GetTransaction(transactionId int) (model.Transaction, error)
	CreateTransaction(transaction model.Transaction) (int64, error)
	UpdateTransaction(transaction model.Transaction) (int64, error)
	DeleteTransaction(id int) error
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
		return -1, err
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return insertedId, nil
}

func (r *transactionRepository) UpdateTransaction(transaction model.Transaction) (int64, error) {
	query := "UPDATE transaction SET Amount=?, Description=? WHERE ID=?"
	result, err := r.db.Exec(
		query,
		transaction.Amount,
		transaction.Description,
		transaction.ID,
	)

	if err != nil {
		return -1, errs.Generic400Err
	}

	updatedId, err := result.RowsAffected()

	if err != nil {
		return -1, errs.Generic400Err
	}

	return updatedId, nil
}

func (r *transactionRepository) DeleteTransaction(id int) error {
	query := "DELETE FROM transaction WHERE ID = ?"

	if _, err := r.db.Exec(query, id); err != nil {
		return errs.Generic400Err
	}

	return nil
}

func (r *transactionRepository) GetTransaction(id int) (model.Transaction, error) {
	var transaction model.Transaction
	query := "SELECT ID, Uuid, Description, Amount, Created, BudgetID, UserID FROM transaction WHERE ID = ?"
	err := r.db.
		QueryRow(query, id).
		Scan(
			&transaction.ID,
			&transaction.Uuid,
			&transaction.Description,
			&transaction.Amount,
			&transaction.Created,
			&transaction.BudgetID,
			&transaction.UserID,
		)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransactionsForBudget(budgetId int) ([]model.Transaction, error) {
	query := "SELECT ID, Uuid, Description, Amount, Created, BudgetID, UserID FROM transaction WHERE BudgetID = ?"
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
