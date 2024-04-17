package budget_handler

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) GetBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	id, err := utils.GetIdFromPath(r.URL.Path)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	budget, err := h.repo.GetBudget(id)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.BudgetNotFound, http.StatusNotFound)
		return
	}

	transactions, err := h.transactionsRepo.GetTransactionsForBudget(id)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusNotFound)
	}

	budgetWithTransaction := model.BudgetWithTransactions{
		Budget:       budget,
		Transactions: transactions,
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgetWithTransaction)
}
