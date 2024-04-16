package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	"tracker/pkg/utils"
)

type TransactionHandler struct {
	repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) TransactionHandler {
	return TransactionHandler{repo}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	budgetId, err := strconv.Atoi(r.URL.Query().Get("budgetId"))

	if err != nil {
		utils.LogError(err.Error())
	}

  // TODO Errors
	transactions, _ := h.repo.GetTransactionsForBudget(budgetId)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(transactions)
}

func (h *TransactionHandler) CreateTransation(w http.ResponseWriter, r *model.ProtectedRequest) {
	var transaction model.Transaction
	budgetId, _ := strconv.Atoi(r.URL.Query().Get("budgetId"))
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&transaction); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if budgetId < 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidBudgetId, http.StatusBadRequest)
		return
	}

	if transaction.Amount == 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidAmount, http.StatusUnprocessableEntity)
		return
	}

	transaction.UserID = r.UserID
	transaction.BudgetID = budgetId
	_, createErr := h.repo.CreateTransaction(transaction)

	if createErr != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Transaction created"}
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	// encoder := json.NewEncoder(w)
	// parts := strings.Split(r.URL.Path, "/")
	// id, err := strconv.Atoi(parts[len(parts)-1])
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	// encoder := json.NewEncoder(w)
	// parts := strings.Split(r.URL.Path, "/")
	// id, err := strconv.Atoi(parts[len(parts)-1])
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	// encoder := json.NewEncoder(w)
	// parts := strings.Split(r.URL.Path, "/")
	// id, err := strconv.Atoi(parts[len(parts)-1])
}
