package budget_handler

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) GetBudgets(w http.ResponseWriter, r *model.ProtectedRequest) {
	budgets, err := h.repo.GetBudgets()
	encoder := json.NewEncoder(w)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(budgets)
}
