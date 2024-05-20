package middleware

import (
	"net/http"
	"tracker/pkg/model"
	"tracker/pkg/repository"
)

type ProtectedHandler func(w http.ResponseWriter, r *model.ProtectedRequest)

func NewAuthMiddleware(repo repository.AuthRepository) authMiddleware {
	return authMiddleware{repo}
}

type authMiddleware struct {
	repo repository.AuthRepository
}

func (m *authMiddleware) ProtectedView(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionid")

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		model.SessionMux.Lock()
		session, ok := model.Sessions[cookie.Value]

		// TODO: check is expired
		if !ok || session.Id == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		model.SessionMux.Unlock()
		next(w, r)
	})
}
