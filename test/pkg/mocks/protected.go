package mocks_test

import (
	"net/http"
	"tracker/pkg/middleware"
	"tracker/pkg/model"
)

func MockProtected(next middleware.ProtectedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		protectedRequest := model.ProtectedRequest{
			Request:   r,
			UserEmail: "mock@mock.com",
			UserID:    1,
		}
		next(w, &protectedRequest)
	})
}
