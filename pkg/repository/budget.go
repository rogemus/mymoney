package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type BudgetRepository interface {
	GetBudget(id int) (model.Budget, error)
	GetBudgets() ([]model.Budget, error)
	CreateBudget(budget model.Budget) error
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
	query := "SELECT id, uuid, created, description, title, userid FROM budgets WHERE id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&b.ID, &b.Uuid, &b.Created, &b.Description, &b.Title, &b.UserID)

	if err == sql.ErrNoRows {
		return b, errs.Budget404Err
	}

	if err != nil {
		return b, errs.Generic400Err
	}

	return b, nil
}

func (r *budgetRepository) GetBudgets() ([]model.Budget, error) {
	query := "SELECT id, uuid, created, description, title, userid FROM budgets"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, errs.Generic400Err
	}

	defer rows.Close()
	budgets := []model.Budget{}

	for rows.Next() {
		var b model.Budget

		if err := rows.Scan(&b.ID, &b.Uuid, &b.Created, &b.Description, &b.Title, &b.UserID); err != nil {
			return nil, errs.Generic400Err
		}

		budgets = append(budgets, b)
	}

	if err := rows.Err(); err != nil {
		return nil, errs.Generic400Err
	}

	return budgets, nil
}

func (r *budgetRepository) CreateBudget(budget model.Budget) error {
	query := "INSERT INTO budgets (title, description, userid) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(
		query,
		budget.Title,
		budget.Description,
		budget.UserID,
	)

	if err != nil {
		return errs.Generic400Err
	}

	return nil
}

func (r *budgetRepository) DeleteBudget(id int) error {
	query := "DELETE FROM budgets WHERE id = $1"

	if _, err := r.db.Exec(query, id); err != nil {
		return errs.Generic400Err
	}

	return nil
}

func (r *budgetRepository) UpdateBudget(budget model.Budget, id int) error {
	query := "UPDATE budgets SET title=$1, description=$2 WHERE id = $3"
	_, err := r.db.Exec(query, budget.Title, budget.Description, id)

	if err != nil {
		return errs.Generic400Err
	}

	return nil
}
