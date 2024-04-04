package repository

import (
	"database/sql"
	"fmt"
	"tracker/pkg/models"
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
	query := "SELECT ID, Uuid, Title, Created, Description FROM budget WHERE ID = ?"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&b.ID, &b.Uuid, &b.Title, &b.Created, &b.Description); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("GetBudget %d: no such budget", id)
		}

		return b, fmt.Errorf("GetBudget %d: %v", id, err)
	}

	return b, nil
}

func (r *budgetRepository) GetBudgets() ([]models.Budget, error) {
	var budgets []models.Budget
	query := "SELECT ID, Uuid, Title, Created, Description FROM budget"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("GetBudgets: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var b models.Budget

		if err := rows.Scan(&b.ID, &b.Uuid, &b.Title, &b.Created, &b.Description); err != nil {
			return nil, fmt.Errorf("GetBudgets: %v", err)
		}

		budgets = append(budgets, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBudgets: %v", err)
	}

	return budgets, nil
}

func (r *budgetRepository) CreateBudget(budget models.Budget) (int64, error) {
	query := "INSERT INTO budget (Title, Description) VALUES (?, ?);"
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
	query := "DELETE from budget WHERE ID = ?;"
	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("DeleteBudget(%d): %v", id, err)
	}

	return nil
}

func (r *budgetRepository) UpdateBudget(budget models.Budget, id int) error {
	query := "UPDATE budget SET Title=?, Description=? WHERE ID = ?;"
	_, err := r.db.Exec(query, budget.Title, budget.Description, id)

	if err != nil {
		return fmt.Errorf("UpdateBudget(%d): %v", id, err)
	}

	return nil
}
