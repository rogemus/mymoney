package models

import "time"

type Transaction struct {
	ID       int
	Uuid     string
	Name     string
	Amount   float32
	Created  time.Time
	BudgetId int
}
