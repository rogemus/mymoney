package models

import (
	"fmt"
)

type Database struct {
  Budgets []Budget
  Transactions []Transaction
}

func (db *Database) Connect() {
  fmt.Println("connecting to db...")
}

func (db *Database) Close() {
  fmt.Println("Disconnecting from db...")
}

