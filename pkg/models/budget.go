package models

import (
	"time"
)

type Budget struct {
	BudgetID    int
	BudgetUuid  string
	Created     time.Time
	Description string
	Title       string
}

type BudgetWithTransactions struct {
	Budget
	Transactions []Transaction
}
