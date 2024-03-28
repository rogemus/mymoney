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
	id, _ := strconv.Atoi(r.PathValue("id"))
	db := database.GetDB()
	br := database.NewBudgetRepository(db)
	budget, err := br.GetBudget(id)

	// TODO handle different type of error
	// TODO write tests
	if err != nil {
		errMsg := err.Error()
		errPayload := models.ErrorPayload{Msg: errMsg}
		utils.LogError(errMsg)
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
		errMsg := err.Error()
		errPayload := models.ErrorPayload{Msg: errMsg}
		utils.LogError(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errPayload)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budgets)
}
