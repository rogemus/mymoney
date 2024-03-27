package web

import (
	"encoding/json"
	"net/http"
)

type ResBody struct {
	Data string
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hello := ResBody{Data: "hello"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hello)
}

func (a *App) GetBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	budget := a.DB.GetBudget(1)
	json.NewEncoder(w).Encode(budget)
}

func (a *App) GetBudgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	budgets := a.DB.GetBudgets()
	json.NewEncoder(w).Encode(budgets)
}
