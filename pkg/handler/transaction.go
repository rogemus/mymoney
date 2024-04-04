package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/repository"
)

type TransactionHandler struct {
	repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) TransactionHandler {
	return TransactionHandler{repo}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	// TODO: HANDLE ERROR
	budgetId, _ := strconv.Atoi(r.URL.Query().Get("budgetId"))
	transactions, _ := h.repo.GetTransactionsForBudget(budgetId)
	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(transactions)
}
