package model

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	Uuid        string    `json:"uuid"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Created     time.Time `json:"created"`
	BudgetID    int       `json:"budgetId"`
	UserID      int       `json:"userId"`
}
