package handler

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/errors"
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
		errors.ErrorResponse(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

  // TODO: Make something nicer
	if user.Email == "" || user.Username == "" || user.Password == "" {
		errors.ErrorResponse(w, errors.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	hashPass, _ := authService.HashPass(user.Password)
	user.Password = hashPass
	_, createUserErr := h.repo.CreateUser(user)

	if createUserErr != nil {
		errors.ErrorResponse(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	payload := model.GenericPayload{Msg: "User created"}
	encoder.Encode(payload)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userReq model.User
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userReq); err != nil {
		errors.ErrorResponse(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

  // TODO: Make something nicer
	if userReq.Password == "" || userReq.Email == "" {
		errors.ErrorResponse(w, errors.Generic422Err, http.StatusUnprocessableEntity)
		return
	}

	userDB, getUserErr := h.repo.GetUserByEmail(userReq.Email)

	if getUserErr != nil {
		errors.ErrorResponse(w, errors.User404Err, http.StatusNotFound)
		return
	}

	validator := userService.UserValidator{User: userReq}

	if !validator.IsEmailValid() || !validator.IsPassValid(userDB.Password) {
		errors.ErrorResponse(w, errors.AuthIvalidPass, http.StatusUnauthorized)
		return
	}

	token := authService.GenerateJwt(userReq.Email)
	_, tokenCreateErr := h.authRepo.CreateToken(token.Token, userReq.Email)

	if tokenCreateErr != nil {
		errors.ErrorResponse(w, tokenCreateErr, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	payload := model.Authenticated{Token: token.Token}
	encoder.Encode(payload)
}
