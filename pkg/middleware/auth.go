package middleware

import (
	"net/http"
	"os"
	"strings"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	"tracker/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type ProtectedHandler func(w http.ResponseWriter, r *model.ProtectedRequest)

func NewAuthMiddleware(repo repository.AuthRepository) authMiddleware {
	return authMiddleware{repo}
}

type authMiddleware struct {
	repo repository.AuthRepository
}

func (m *authMiddleware) ProtectedRoute(next ProtectedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if len(bearerToken) == 0 {
			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		reqToken, err := splitHeader(bearerToken)
		if err != nil {
			utils.LogError(err.Error())
			errs.ErrorResponse(w, err, http.StatusUnauthorized)
			return
		}

		claims := &model.Claims{}
		tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.LogError(err.Error())
				errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
				return
			}

			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		if !tkn.Valid {
			utils.LogError(err.Error())
			errs.ErrorResponse(w, errs.Generic401Err, http.StatusUnauthorized)
			return
		}

		protectedRequest := model.ProtectedRequest{
			Request:   r,
			UserEmail: claims.UserEmail,
			UserID:    claims.UserID,
		}
		w.Header().Set("Content-Type", "application/json")
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
