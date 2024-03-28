package middleware

import (
	"fmt"
	"net/http"
	"tracker/pkg/utils"
)

func LogReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logPattern := `%s - "%s %s %s"`
		utils.LogInfo(fmt.Sprintf(logPattern, r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))
		next.ServeHTTP(w, r)
	})
}
