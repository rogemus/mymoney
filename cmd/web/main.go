package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"tracker/pkg/handler"
	"tracker/pkg/middleware"
	"tracker/pkg/repository"
	"tracker/pkg/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	// Get .env file
	if err = godotenv.Load(); err != nil {
		utils.LogFatal("Error loading .env file")
	}

	// Init DB
	cfg := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Addr:      os.Getenv("DB_ADDR"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err = db.Ping(); err != nil {
		utils.LogFatal(err.Error())
	}
	utils.LogInfo("Connected do Database 🎉 !!!")

	// Init Routes
	mux := http.NewServeMux()

	// API: Transactions
	transactionRepo := repository.NewTransactionRepository(db)
	transactionHandler := handler.NewTransactionHandler(transactionRepo)
	mux.HandleFunc("GET /transactions", transactionHandler.GetTransactions)

	// API: Budget
	budgetRepo := repository.NewBudgetRepository(db)
	budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

	mux.HandleFunc("GET /budget/{id}", budgetHandler.GetBudget)
	mux.HandleFunc("DELETE /budget/{id}", budgetHandler.DeleteBudget)
	mux.HandleFunc("PUT /budget/{id}", budgetHandler.UpdateBudget)
	mux.HandleFunc("GET /budgets", budgetHandler.GetBudgets)
	mux.HandleFunc("POST /budgets", budgetHandler.CreateBudget)

  // API: Token
  authRepo := repository.NewAuthRepository(db)

	// API: User
	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo, authRepo)

	mux.HandleFunc("POST /register", userHandler.RegisterUser)
	mux.HandleFunc("POST /login", userHandler.LoginUser)


	// API: Public Files
	publicFiles := http.FileServer(http.Dir("./ui/public"))
	mux.Handle("/", publicFiles)

	routes := middleware.LogReq(middleware.ServeJson(mux))

	// Init server
	srv := &http.Server{
		Addr:    ":3333",
		Handler: routes,
	}

	// Start Server
	utils.LogInfo(fmt.Sprintf("Listening on port: %v ...", ":3333"))
	srv.ListenAndServe()
}
