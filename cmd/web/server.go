package web

import (
	"fmt"
	"net/http"
	"tracker/pkg/handlers"
	"tracker/pkg/middleware"
	"tracker/pkg/utils"
)

type App struct {
	Addr      string
	PublicDir string
}

func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()

	// API: Budget
	mux.HandleFunc("GET /budget/{id}", handlers.GetBudget)
	mux.HandleFunc("DELETE /budget/{id}", handlers.DeleteBudget)
	mux.HandleFunc("PUT /budget/{id}", handlers.UpdateBudget)
	mux.HandleFunc("GET /budgets", handlers.GetBudgets)
	mux.HandleFunc("POST /budgets", handlers.CreateBudget)

	publicFiles := http.FileServer(http.Dir(a.PublicDir))
	mux.Handle("/", publicFiles)

	return middleware.LogReq(mux)
}

func (a *App) RunServer() {
	srv := &http.Server{
		Addr:    a.Addr,
		Handler: a.Routes(),
	}

	utils.LogInfo(fmt.Sprintf("Listening on port: %v ...", a.Addr))
	err := srv.ListenAndServe()

	if err != nil {
		utils.LogFatal("Error while oppening the server :()")
	}
}
