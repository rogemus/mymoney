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

func (h *BudgetHandler) CreateTransation(w http.ResponseWriter, r *model.ProtectedRequest) {
	var transaction model.Transaction
	budgetId, err := strconv.Atoi(r.URL.Query().Get("budgetId"))
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err != nil || budgetId == 0 {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.BudgetInvalidBudgetId, http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&transaction); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	// TODO: Move to service
	if transaction.Amount <= 0 {
		utils.LogError(errs.TransactionInvalidAmount.Error())
		errs.ErrorResponse(w, errs.TransactionInvalidAmount, http.StatusUnprocessableEntity)
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

	transaction.UserID = r.UserID
	transaction.BudgetID = budgetId

	if _, err := h.transactionsRepo.CreateTransaction(transaction); err != nil {
		utils.LogError(err.Error())
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Transaction created"}
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(payload)
}


