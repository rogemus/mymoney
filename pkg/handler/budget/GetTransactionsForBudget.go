package budget_handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) GetTransactions(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	budgetId, err := strconv.Atoi(r.URL.Query().Get("budgetId"))

	if err != nil || budgetId == 0 {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.BudgetInvalidBudgetId, http.StatusBadRequest)
		return
	}

	_, err = h.repo.GetBudget(budgetId)
	if err == sql.ErrNoRows {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.BudgetNotFound, http.StatusNotFound)
		return
	}

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
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
