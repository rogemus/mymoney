package main

import (
	"fmt"
	"net/http"
	"os"
	"tracker/pkg/database"
	"tracker/pkg/handlers"
	"tracker/pkg/middleware"
	"tracker/pkg/utils"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	return middleware.LogReq(middleware.ServeJson(mux))
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

func main() {
	if err := godotenv.Load(); err != nil {
		utils.LogFatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Addr:      "localhost:3309",
		DBName:    "tracker",
		ParseTime: true,
	}

	err := database.InitDB(cfg.FormatDSN())

	if err != nil {
		utils.LogFatal(err.Error())
	}

	app := &App{
		Addr:      ":3333",
		PublicDir: "./ui/public/",
	}

	app.RunServer()
}
