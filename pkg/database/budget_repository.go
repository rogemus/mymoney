package database

import (
	"database/sql"
	"fmt"
	"tracker/pkg/models"
	"tracker/pkg/utils"
)

type BudgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) *BudgetRepository {
	return &BudgetRepository{db}
}

func (br *BudgetRepository) GetBudget(id int) (models.Budget, error) {
	var b models.Budget
	query := "SELECT ID, Uuid, Title, Created, Description FROM budget WHERE ID = ?"
	utils.LogInfo(fmt.Sprintf("Looking for Budget(%b)...", id))
	row := br.db.QueryRow(query, id)

	if err := row.Scan(&b.ID, &b.Uuid, &b.Title, &b.Created, &b.Description); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("GetBudget %d: no such budget", id)
		}

		return b, fmt.Errorf("GetBudget %d: %v", id, err)
	}

	utils.LogInfo(fmt.Sprintf("Found: Budget(%d, %s)", b.ID, b.Title))
	return b, nil
}

func (br *BudgetRepository) GetBudgets() ([]models.Budget, error) {
	var budgets []models.Budget
	query := "SELECT ID, Uuid, Title, Created, Description FROM budget"
	utils.LogInfo("Looking for []Budget...")
	rows, err := br.db.Query(query)

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

	utils.LogInfo("Found: []Budget")
	return budgets, nil
}

func (br *BudgetRepository) CreateBudget(budget models.Budget) (int64, error) {
	query := "INSERT INTO budget (Title, Description) VALUES (?, ?);"
	result, err := br.db.Exec(query, budget.Title, budget.Description)

	if err != nil {
		return 0, fmt.Errorf("CreateBudget: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("CreateBudget: %v", err)
	}

	return id, nil
}
