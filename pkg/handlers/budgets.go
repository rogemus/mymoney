package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/database"
	"tracker/pkg/models"
	"tracker/pkg/utils"
)

func GetBudget(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	tr := database.NewTransactionRepository(db)

	id, _ := strconv.Atoi(r.PathValue("id"))
	budget, err := br.GetBudget(id)
	transactions, _ := tr.GetTransactions(id)
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

func GetBudgets(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	budgets, err := br.GetBudgets()
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

func CreateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
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
	br.CreateBudget(budget)
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}

func DeleteBudget(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	id, _ := strconv.Atoi(r.PathValue("id"))
	encoder := json.NewEncoder(w)

	if err := br.DeleteBudget(id); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
		return
	}

	payload := models.GenericPayload{Msg: "Budget deleted"}
	w.WriteHeader(http.StatusAccepted)
	encoder.Encode(payload)
}

func UpdateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	id, _ := strconv.Atoi(r.PathValue("id"))
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errPayload)
		return
	}

	payload := models.GenericPayload{Msg: "Budget updated"}
	br.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}
