package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/models"
	"tracker/pkg/utils"
)

func (a *App) GetBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.PathValue("id"))
	budget, err := a.Database.GetBudget(id)

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

func (a *App) GetBudgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	budgets, err := a.Database.GetBudgets()

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
