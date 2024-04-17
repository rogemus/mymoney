package budget_handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/utils"
)

func (h *BudgetHandler) DeleteBudget(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	id, err := utils.GetIdFromPath(r.URL.Path)

	if err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	_, err = h.repo.GetBudget(id)

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

	if err := h.repo.DeleteBudget(id); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, 400)
		return
	}

	payload := model.GenericPayload{Msg: "Budget deleted"}
	w.WriteHeader(http.StatusNoContent)
	encoder.Encode(payload)
}
