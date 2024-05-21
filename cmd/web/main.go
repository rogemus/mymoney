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

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var err error

	// Get .env file
	if err = godotenv.Load(); err != nil {
		utils.LogFatal("Error loading .env file")
	}

	// Init DB
	cfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", cfg)

	if err = db.Ping(); err != nil {
		utils.LogFatal(err.Error())
	}
	utils.LogInfo("Connected do Database 🎉 !!!")

	// Init Routes
	mux := http.NewServeMux()

	// API: Token
	authRepo := repository.NewAuthRepository(db)
	authMiddleware := middleware.NewAuthMiddleware(authRepo)
	protected := authMiddleware.ProtectedView

	// User
	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo, authRepo)

	mux.HandleFunc("GET /", userHandler.LoginView)
	mux.HandleFunc("GET /register", userHandler.RegisterView)
	mux.HandleFunc("GET /logout", userHandler.LogoutView)
	mux.HandleFunc("POST /signup", userHandler.Signup)
	mux.HandleFunc("POST /signin", userHandler.Signin)
	mux.HandleFunc("POST /signout", userHandler.Signout)

	// Dashboard
	dashboardHandler := handler.NewDashboardHandler()

	mux.HandleFunc("GET /dashboard", protected(dashboardHandler.MainView))

	// API: Public Files
	publicFiles := http.FileServer(http.Dir("ui/public/browser"))
	mux.Handle("/", publicFiles)

	routes := middleware.LogReq(mux)
	addr := fmt.Sprintf("%s:%s", os.Getenv("ADDR"), os.Getenv("PORT"))

	// Init server
	srv := &http.Server{
		Addr:    addr,
		Handler: routes,
	}

	// Start Server
	utils.LogInfo(fmt.Sprintf("Listening on: %v ...", addr))
	servErr := srv.ListenAndServe()

	if servErr != nil {
		utils.LogFatal(servErr.Error())
	}
}
