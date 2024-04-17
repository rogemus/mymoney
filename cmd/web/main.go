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

	// API: Token
	authRepo := repository.NewAuthRepository(db)
	authMiddleware := middleware.NewAuthMiddleware(authRepo)
	protected := authMiddleware.ProtectedRoute

	// API: Transactions
	transactionRepo := repository.NewTransactionRepository(db)
	transactionHandler := handler.NewTransactionHandler(transactionRepo)

	mux.HandleFunc("POST /transactions", protected(transactionHandler.CreateTransation))
	mux.HandleFunc("DELETE /transactions/{id}", protected(transactionHandler.DeleteTransaction))
	mux.HandleFunc("POST /transactions/{id}", protected(transactionHandler.UpdateTransaction))
	mux.HandleFunc("GET /transactions/{id}", protected(transactionHandler.GetTransaction))

	// API: Budget
	budgetRepo := repository.NewBudgetRepository(db)
	budgetHandler := handler.NewBudgetHandler(budgetRepo, transactionRepo)

	mux.HandleFunc("GET /budget/{id}", protected(budgetHandler.GetBudget))
	mux.HandleFunc("DELETE /budget/{id}", protected(budgetHandler.DeleteBudget))
	mux.HandleFunc("PUT /budget/{id}", protected(budgetHandler.UpdateBudget))
	mux.HandleFunc("GET /budgets", protected(budgetHandler.GetBudgets))
	mux.HandleFunc("POST /budgets", protected(budgetHandler.CreateBudget))

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
	servErr := srv.ListenAndServe()

	if servErr != nil {
		utils.LogFatal(servErr.Error())
	}
}
