package utils

import (
	"tracker/pkg/models"
)

func ErrRes(err error) models.GenericPayload {
	errMsg := err.Error()
	LogError(errMsg)
	return models.GenericPayload{Msg: errMsg}
}
