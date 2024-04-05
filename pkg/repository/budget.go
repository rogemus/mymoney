package repository

import (
	"database/sql"
	"fmt"
	"tracker/pkg/models"
	errors "tracker/pkg/utils"
)

type BudgetRepository interface {
	GetBudget(id int) (models.Budget, error)
	GetBudgets() ([]models.Budget, error)
	CreateBudget(budget models.Budget) (int64, error)
	UpdateBudget(budget models.Budget, id int) error
	DeleteBudget(id int) error
}

type budgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) BudgetRepository {
	return &budgetRepository{db}
}

func (r *budgetRepository) GetBudget(id int) (models.Budget, error) {
	var b models.Budget
	query := "SELECT ID, Uuid, Created, Description, Title FROM budget WHERE ID = ?"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&b.ID, &b.Uuid, &b.Created, &b.Description, &b.Title)

	switch err {
	case sql.ErrNoRows:
		return b, errors.Budget404Err
	case nil:
		return b, nil
	default:
		return b, errors.Generic400Err
	}
}

func (r *budgetRepository) GetBudgets() ([]models.Budget, error) {
	query := "SELECT ID, Uuid, Created, Description, Title FROM budget"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, errors.Generic400Err
	}

	defer rows.Close()
	budgets := []models.Budget{}

	for rows.Next() {
		var b models.Budget

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

func (r *budgetRepository) CreateBudget(budget models.Budget) (int64, error) {
	query := "INSERT INTO budget (Title, Description) VALUES (?, ?)"
	result, err := r.db.Exec(query, budget.Title, budget.Description)

	if err != nil {
		return 0, fmt.Errorf("CreateBudget: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("CreateBudget: %v", err)
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

func (r *budgetRepository) UpdateBudget(budget models.Budget, id int) error {
	query := "UPDATE budget SET Title=?, Description=? WHERE ID = ?"
	_, err := r.db.Exec(query, budget.Title, budget.Description, id)

	if err != nil {
		return fmt.Errorf("UpdateBudget(%d): %v", id, err)
	}

	return nil
}
