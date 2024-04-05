package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	"tracker/pkg/utils"
	errors "tracker/pkg/utils"
)

type BudgetHandler struct {
	repo             repository.BudgetRepository
	transactionsRepo repository.TransactionRepository
}

func NewBudgetHandler(repo repository.BudgetRepository, transactionsRepo repository.TransactionRepository) BudgetHandler {
	return BudgetHandler{repo, transactionsRepo}
}

func (h *BudgetHandler) GetBudget(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	budget, err := h.repo.GetBudget(id)

	if err == errors.Budget404Err {
		utils.ErrRes(w, errors.Budget404Err, http.StatusNotFound)
		return
	}

	transactions, _ := h.transactionsRepo.GetTransactionsForBudget(id)
	budgetWithTransaction := models.BudgetWithTransactions{
		Budget:       budget,
		Transactions: transactions,
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgetWithTransaction)
}

func (h *BudgetHandler) GetBudgets(w http.ResponseWriter, r *http.Request) {
	budgets, err := h.repo.GetBudgets()
	encoder := json.NewEncoder(w)

	if err != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgets)
}

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		utils.ErrRes(w, err, 500)
		return
	}

	// TODO handle errors
	// TODO write tests
	payload := models.GenericPayload{Msg: "Budget created"}
	h.repo.CreateBudget(budget)
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}

func (h *BudgetHandler) DeleteBudget(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	encoder := json.NewEncoder(w)

	if err := h.repo.DeleteBudget(id); err != nil {
		utils.ErrRes(w, err, 500)
		return
	}

	payload := models.GenericPayload{Msg: "Budget deleted"}
	w.WriteHeader(http.StatusAccepted)
	encoder.Encode(payload)
}

func (h *BudgetHandler) UpdateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	id, _ := strconv.Atoi(r.PathValue("id"))
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		utils.ErrRes(w, err, 500)
		return
	}

	payload := models.GenericPayload{Msg: "Budget updated"}
	h.repo.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}
