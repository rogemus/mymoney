package web

import "net/http"

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()

	// API
	mux.HandleFunc("GET /hello", GetHello)

  // API: Budget
  mux.HandleFunc("GET /budget/{id}", a.GetBudget)
  mux.HandleFunc("GET /budgets", a.GetBudgets)

	publicFiles := http.FileServer(http.Dir(a.PublicDir))
	mux.Handle("/", publicFiles)

	return mux
}
