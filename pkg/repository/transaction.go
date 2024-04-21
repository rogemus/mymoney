package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type TransactionRepository interface {
	GetTransactionsForBudget(budgetId int) ([]model.Transaction, error)
	GetTransaction(transactionId int) (model.Transaction, error)
	CreateTransaction(transaction model.Transaction) error
	UpdateTransaction(transaction model.Transaction) (int64, error)
	DeleteTransaction(id int) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction model.Transaction) error {
	query := "INSERT INTO transactions (description, amount, budgetid, userid) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(
		query,
		transaction.Description,
		transaction.Amount,
		transaction.BudgetID,
		transaction.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) UpdateTransaction(transaction model.Transaction) (int64, error) {
	query := "UPDATE transactions SET amount=$1, description=$2 WHERE id=$3"
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
	query := "DELETE FROM transactions WHERE id = $1"

	if _, err := r.db.Exec(query, id); err != nil {
		return errs.Generic400Err
	}

	return nil
}

func (r *transactionRepository) GetTransaction(id int) (model.Transaction, error) {
	var transaction model.Transaction
	query := "SELECT id, uuid, description, amount, created, budgetid, userid FROM transactions WHERE id = $1"
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

	if err == sql.ErrNoRows {
		return transaction, errs.Transaction404Err
	}

	if err != nil {
		return transaction, errs.Generic400Err
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransactionsForBudget(budgetId int) ([]model.Transaction, error) {
	query := "SELECT id, uuid, description, amount, created, budgetid, userid FROM transactions WHERE budgetid = $1"
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
