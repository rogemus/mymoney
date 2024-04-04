package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/models"
	"tracker/pkg/repository"
	"tracker/pkg/utils"
)

type BudgetHandler struct {
	repo             repository.BudgetRepository
	transactionsRepo repository.TransactionRepository
}

func NewBudgetHandler(repo repository.BudgetRepository, transactionsRepo repository.TransactionRepository) BudgetHandler {
	return BudgetHandler{repo, transactionsRepo}
}

func (h *BudgetHandler) GetBudget(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	budget, err := h.repo.GetBudget(id)
	transactions, _ := h.transactionsRepo.GetTransactionsForBudget(id)
	encoder := json.NewEncoder(w)
	budgetWithTransaction := models.BudgetWithTransactions{
		Budget:       budget,
		Transactions: transactions,
	}

	// TODO handle different type of error
	// TODO write tests
	if err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgetWithTransaction)
}

func (h *BudgetHandler) GetBudgets(w http.ResponseWriter, r *http.Request) {
	budgets, err := h.repo.GetBudgets()
	encoder := json.NewEncoder(w)

	// TODO handle different type of error
	// TODO write tests
	if err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgets)
}

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
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
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
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
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
		return
	}

	payload := models.GenericPayload{Msg: "Budget updated"}
	h.repo.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}
