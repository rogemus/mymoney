package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"tracker/pkg/database"
	"tracker/pkg/handlers"
	"tracker/pkg/middleware"
	"tracker/pkg/utils"
)

type App struct {
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

func (a *App) RunServer(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Routes(),
	}

	utils.LogInfo(fmt.Sprintf("Listening on port: %v ...", addr))
	err := srv.ListenAndServe()

	if err != nil {
		utils.LogFatal("Error while oppening the server :()")
	}
}

func (a App) Initialize(
	dbUser string,
	dbPass string,
	dbAddr string,
	dbName string,
) {
	if err := godotenv.Load(); err != nil {
		utils.LogFatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:      dbUser,
		Passwd:    dbPass,
		Addr:      dbAddr,
		DBName:    dbName,
		ParseTime: true,
	}

	err := database.InitDB(cfg.FormatDSN())

	if err != nil {
		utils.LogFatal(err.Error())
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		utils.LogFatal("Error loading .env file")
	}

	app := &App{
		PublicDir: "./ui/public/",
	}

	app.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_NAME"),
	)

	app.RunServer(":3333")
}
