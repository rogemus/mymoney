package models

import "time"

type Transaction struct {
	TransactionID   int
	TransactionUuid string
	Description     string
	Amount          float32
	Created         time.Time
	BudgetID        int
}
