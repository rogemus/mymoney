package budget_handler

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	var budget model.Budget
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&budget); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

  // TODO: move to service
	if budget.Title == "" {
		utils.LogError(errs.BudgetInvalidTitle.Error())
		errs.ErrorResponse(w, errs.BudgetInvalidTitle, http.StatusBadRequest)
		return
	}

	budget.UserID = r.UserID
	_, err := h.repo.CreateBudget(budget)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Budget created"}
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}
