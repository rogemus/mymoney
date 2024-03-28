package web

import (
	"net/http"
	"tracker/pkg/handlers"
	"tracker/pkg/middleware"
)

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()

	// API: Budget
	mux.HandleFunc("GET /budget/{id}", handlers.GetBudget)
	mux.HandleFunc("GET /budgets", handlers.GetBudgets)

	publicFiles := http.FileServer(http.Dir(a.PublicDir))
	mux.Handle("/", publicFiles)

	return middleware.LogReq(mux)
}
