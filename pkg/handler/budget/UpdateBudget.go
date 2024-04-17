package budget_handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) UpdateBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	var budget model.Budget
	id, err := utils.GetIdFromPath(r.URL.Path)
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	_, err = h.repo.GetBudget(id)
	if err == sql.ErrNoRows {
		errs.ErrorResponse(w, errs.BudgetNotFound, http.StatusNotFound)
		return
	}

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&budget); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if budget.Title == "" {
		utils.LogError(errs.BudgetInvalidTitle.Error())
		errs.ErrorResponse(w, errs.BudgetInvalidTitle, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Budget updated"}
	h.repo.UpdateBudget(budget, id)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}
