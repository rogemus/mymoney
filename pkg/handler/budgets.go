package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
)

type BudgetHandler struct {
	repo             repository.BudgetRepository
	transactionsRepo repository.TransactionRepository
}

func NewBudgetHandler(repo repository.BudgetRepository, transactionsRepo repository.TransactionRepository) BudgetHandler {
	return BudgetHandler{repo, transactionsRepo}
}

func (h *BudgetHandler) GetBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	budget, err := h.repo.GetBudget(id)

	if err != nil {
		errs.ErrorResponse(w, errs.Budget404Err, http.StatusNotFound)
		return
	}

	// TODO: Errors
	transactions, _ := h.transactionsRepo.GetTransactionsForBudget(id)

	budgetWithTransaction := model.BudgetWithTransactions{
		Budget:       budget,
		Transactions: transactions,
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgetWithTransaction)
}

func (h *BudgetHandler) GetBudgets(w http.ResponseWriter, r *model.ProtectedRequest) {
	budgets, err := h.repo.GetBudgets()
	encoder := json.NewEncoder(w)

	if err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgets)
}

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	var budget model.Budget
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	// move to service
	if budget.Title == "" {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	budget.UserID = r.UserID
	payload := model.GenericPayload{Msg: "Budget created"}
	h.repo.CreateBudget(budget)
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}

func (h *BudgetHandler) DeleteBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if _, err := h.repo.GetBudget(id); err == errs.Budget404Err {
		errs.ErrorResponse(w, errs.Budget404Err, http.StatusNotFound)
		return
	}

	if err := h.repo.DeleteBudget(id); err != nil {
		errs.ErrorResponse(w, err, 400)
		return
	}

	payload := model.GenericPayload{Msg: "Budget deleted"}
	w.WriteHeader(http.StatusNoContent)
	encoder.Encode(payload)
}

func (h *BudgetHandler) UpdateBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	var budget model.Budget
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	encoder := json.NewEncoder(w)

	if err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if _, err := h.repo.GetBudget(id); err == errs.Budget404Err {
		errs.ErrorResponse(w, errs.Budget404Err, http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	// Move to service
	if budget.Title == "" {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Budget updated"}
	h.repo.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}

func (h *BudgetHandler) CreateTransation(w http.ResponseWriter, r *model.ProtectedRequest) {
	var transaction model.Transaction
	budgetId, err := strconv.Atoi(r.URL.Query().Get("budgetId"))
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err != nil || budgetId == 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidBudgetId, http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&transaction); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidAmount, http.StatusUnprocessableEntity)
		return
	}

	if _, err := h.repo.GetBudget(budgetId); err != nil {
		errs.ErrorResponse(w, errs.Budget404Err, http.StatusNotFound)
		return
	}

	transaction.UserID = r.UserID
	transaction.BudgetID = budgetId

	if err := h.transactionsRepo.CreateTransaction(transaction); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Transaction created"}
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}

func (h *BudgetHandler) GetTransactions(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	budgetId, err := strconv.Atoi(r.URL.Query().Get("budgetId"))

	if err != nil || budgetId == 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidBudgetId, http.StatusBadRequest)
		return
	}

	if _, err := h.repo.GetBudget(budgetId); err != nil {
		errs.ErrorResponse(w, errs.Budget404Err, http.StatusNotFound)
		return
	}

	transactions, err := h.transactionsRepo.GetTransactionsForBudget(budgetId)

	if err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(transactions)
}
