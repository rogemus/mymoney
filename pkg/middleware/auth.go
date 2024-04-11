package middleware

import (
	"net/http"
	"os"
	"strings"
	"tracker/pkg/errs"
	"tracker/pkg/model"

	"github.com/golang-jwt/jwt/v5"
)

type protectedHandler func(w http.ResponseWriter, r *model.ProtectedRequest)

func Protected(next protectedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		reqToken, err := splitHeader(bearerToken)

		if err != nil {
			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		claims := &model.Claims{}

    // TODO: Check if token in db
    // TODO: Refactor this
    tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
				return
			}

			errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		protectedRequest := model.ProtectedRequest{Request: r, UserEmail: claims.UserEmail}
		next(w, &protectedRequest)
	})
}

func splitHeader(header string) (string, error) {
	parts := strings.Split(header, " ")

	if len(parts) != 2 {
		return "", errs.AuthInvalidHeader
	}

	return parts[1], nil
}
