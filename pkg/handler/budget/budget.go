package budget_handler

import "tracker/pkg/repository"

type BudgetHandler struct {
	repo             repository.BudgetRepository
	transactionsRepo repository.TransactionRepository
}

func NewBudgetHandler(
	repo repository.BudgetRepository,
	transactionsRepo repository.TransactionRepository,
) BudgetHandler {
	return BudgetHandler{repo, transactionsRepo}
}
