package handler

import (
	"encoding/json"
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

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	// TODO: Make something nicer
	if user.Email == "" || user.Username == "" || user.Password == "" {
		errs.ErrorResponse(w, errs.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	hashPass, _ := authService.HashPass(user.Password)
	user.Password = hashPass
	createUserErr := h.repo.CreateUser(user)

	if createUserErr != nil {
		errs.ErrorResponse(w, errs.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	payload := model.GenericPayload{Msg: "User created"}
	encoder.Encode(payload)
}

func (h *UserHandler) LoginView(w http.ResponseWriter, r *http.Request) {
	// TODO: handle err
	// TODO: redirect if cookie is valid

	templ, _ := template.
		New("login").
		ParseFiles("ui/views/login.html", "ui/views/_base.html")

	templ.ExecuteTemplate(w, "base", nil)
}

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
