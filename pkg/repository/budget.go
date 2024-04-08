package repository

import (
	"database/sql"
	"fmt"
	"tracker/pkg/model"
	errors "tracker/pkg/utils"
)

type BudgetRepository interface {
	GetBudget(id int) (model.Budget, error)
	GetBudgets() ([]model.Budget, error)
	CreateBudget(budget model.Budget) (int64, error)
	UpdateBudget(budget model.Budget, id int) error
	DeleteBudget(id int) error
}

type budgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) BudgetRepository {
	return &budgetRepository{db}
}

func (r *budgetRepository) GetBudget(id int) (model.Budget, error) {
	var b model.Budget
	query := "SELECT ID, Uuid, Created, Description, Title FROM budget WHERE ID = ?"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&b.ID, &b.Uuid, &b.Created, &b.Description, &b.Title)

	if err == sql.ErrNoRows {
		return b, errors.Budget404Err
	}

	if err != nil {
		return b, errors.Generic400Err
	}

	return b, nil
}

func (r *budgetRepository) GetBudgets() ([]model.Budget, error) {
	query := "SELECT ID, Uuid, Created, Description, Title FROM budget"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, errors.Generic400Err
	}

	defer rows.Close()
	budgets := []model.Budget{}

	for rows.Next() {
		var b model.Budget

		if err := rows.Scan(&b.ID, &b.Uuid, &b.Created, &b.Description, &b.Title); err != nil {
			return nil, errors.Generic400Err
		}

		budgets = append(budgets, b)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Generic400Err
	}

	return budgets, nil
}

func (r *budgetRepository) CreateBudget(budget model.Budget) (int64, error) {
	query := "INSERT INTO budget (Title, Description) VALUES (?, ?)"
	result, err := r.db.Exec(query, budget.Title, budget.Description)

	if err != nil {
		return -1, errors.Generic400Err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, errors.Generic400Err
	}

	return id, nil
}

func (r *budgetRepository) DeleteBudget(id int) error {
	query := "DELETE FROM budget WHERE ID = ?"

	if _, err := r.db.Exec(query, id); err != nil {
		return errors.Generic400Err
	}

	return nil
}

func (r *budgetRepository) UpdateBudget(budget model.Budget, id int) error {
	query := "UPDATE budget SET Title=?, Description=? WHERE ID = ?"
	_, err := r.db.Exec(query, budget.Title, budget.Description, id)

	if err != nil {
		return fmt.Errorf("UpdateBudget(%d): %v", id, err)
	}

	return nil
}
