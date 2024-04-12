package model

import (
	"time"
)

type Budget struct {
	ID          int       `json:"id"`
	Uuid        string    `json:"uuid"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	UserID      int       `json:"userId"`
}

type BudgetWithTransactions struct {
	Budget       `json:"budget"`
	Transactions []Transaction `json:"transactions"`
}
