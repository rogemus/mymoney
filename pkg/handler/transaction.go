package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
)

type TransactionHandler struct {
	repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) TransactionHandler {
	return TransactionHandler{repo}
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	var transaction model.Transaction

	if err != nil {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	if err := decoder.Decode(&transaction); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		errs.ErrorResponse(w, errs.TransactionInvalidAmount, http.StatusUnprocessableEntity)
		return
	}

	transactionDb, err := h.repo.GetTransaction(id)
	if err != nil {
		errs.ErrorResponse(w, errs.Transaction404Err, http.StatusNotFound)
		return
	}

	transactionDb.Amount = transaction.Amount
	transactionDb.Description = transaction.Description
	if _, err := h.repo.UpdateTransaction(transactionDb); err != nil {
		errs.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Transaction updated"}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	if _, err := h.repo.GetTransaction(id); err != nil {
		errs.ErrorResponse(w, errs.Transaction404Err, http.StatusNotFound)
		return
	}

	if err := h.repo.DeleteTransaction(id); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	payload := model.GenericPayload{Msg: "Transaction deleted"}
	w.WriteHeader(204)
	encoder.Encode(payload)
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *model.ProtectedRequest) {
	encoder := json.NewEncoder(w)
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	transaction, err := h.repo.GetTransaction(id)

	if err != nil {
		errs.ErrorResponse(w, errs.Transaction404Err, http.StatusNotFound)
		return
	}

	w.WriteHeader(200)
	encoder.Encode(transaction)
}
