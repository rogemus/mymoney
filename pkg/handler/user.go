package handler

import (
	"html/template"
	"net/http"
	"time"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	authService "tracker/pkg/service"
	userService "tracker/pkg/service"
)

type UserHandler struct {
	repo     repository.UserRepository
	authRepo repository.AuthRepository
}

func NewUserHandler(repo repository.UserRepository, authRepo repository.AuthRepository) UserHandler {
	return UserHandler{repo, authRepo}
}

func (h *UserHandler) LogoutView(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.
		New("logout").
		ParseFiles("ui/views/logout.html", "ui/views/_base.html")

	templ.ExecuteTemplate(w, "base", nil)
}

func (h *UserHandler) RegisterView(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.
		New("register").
		ParseFiles("ui/views/register.html", "ui/views/_base.html")

	templ.ExecuteTemplate(w, "base", nil)
}

func (h *UserHandler) LoginView(w http.ResponseWriter, r *http.Request) {
	// TODO: redirect if cookie is valid
	templ, _ := template.
		New("login").
		ParseFiles("ui/views/login.html", "ui/views/_base.html")

	templ.ExecuteTemplate(w, "base", nil)
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	user := model.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// TODO: Make something nicer
	if user.Password == "" || user.Username == "" || user.Email == "" {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	hashPass, _ := authService.HashPass(user.Password)
	user.Password = hashPass

	if err := h.repo.CreateUser(user); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	// TODO: display notification after user is created
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *UserHandler) Signout(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) Signin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	userReq := model.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// TODO: Make something nicer
	if userReq.Password == "" || userReq.Email == "" {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	userDB, getUserErr := h.repo.GetUserByEmail(userReq.Email)

	if getUserErr != nil {
		errs.ErrorResponse(w, errs.User404Err, http.StatusNotFound)
		return
	}

	validator := userService.UserValidator{User: userReq}

	if !validator.IsEmailValid() || !validator.IsPassValid(userDB.Password) {
		errs.ErrorResponse(w, errs.AuthIvalidPass, http.StatusUnauthorized)
		return
	}

	sessionid := authService.GenerateSessionId()
	// TODO: Store session in db
	session := model.Session{
		Id:        sessionid,
		UserEmail: userDB.Email,
		Duration:  int(model.SessionDuration),
		ExpiresAt: time.Now().Add(model.SessionDuration),
	}

	model.SessionMux.Lock()
	model.Sessions[sessionid] = session
	model.SessionMux.Unlock()

	cookie := http.Cookie{
		Name:   "sessionid",
		Value:  sessionid,
		MaxAge: int(model.SessionDuration),
		Path:   "/",
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
