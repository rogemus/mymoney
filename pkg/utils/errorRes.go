package utils

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/models"
)

func ErrRes(w http.ResponseWriter, err error, errorStatus int) {
	encoder := json.NewEncoder(w)
	errMsg := err.Error()
	payload := models.GenericPayload{Msg: errMsg}
	w.WriteHeader(errorStatus)
	encoder.Encode(payload)
}
