package web

import (
	"log"
	"net/http"
)

func LogReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logPattern := `%s - "%s %s %s"`
		log.Printf(logPattern, r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
