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
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	id, _ := strconv.Atoi(r.PathValue("id"))
	budget, err := br.GetBudget(id)

	// TODO handle different type of error
	// TODO write tests
	if err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budget)
}

func GetBudgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	budgets, err := br.GetBudgets()

	// TODO handle different type of error
	// TODO write tests
	if err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budgets)
}

func CreateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	db := database.GetDB()
	br := database.NewBudgetRepository(db)

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	// TODO handle errors
	// TODO write tests
	br.CreateBudget(budget)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}

func DeleteBudget(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	id, _ := strconv.Atoi(r.PathValue("id"))

	if err := br.DeleteBudget(id); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("deleted"))
}

func UpdateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	id, _ := strconv.Atoi(r.PathValue("id"))

	if err := json.NewDecoder(r.Body).Decode(&budget); err != nil {
		errPayload := utils.ErrRes(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	br.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("updated"))
}
