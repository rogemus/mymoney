package utils

import (
	"tracker/pkg/models"
)

func ErrRes(err error) models.ErrorPayload {
	errMsg := err.Error()
	LogError(errMsg)
	return models.ErrorPayload{Msg: errMsg}
}
