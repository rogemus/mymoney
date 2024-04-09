package handler

import (
	"encoding/json"
	"net/http"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	authService "tracker/pkg/service"
	userService "tracker/pkg/service"
	"tracker/pkg/utils"
	errors "tracker/pkg/utils"
)

// TEMP: Move to DB
var tokens = make([]model.Token, 0)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) UserHandler {
	return UserHandler{repo}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	// Move to Service
	if user.Email == "" || user.Username == "" || user.Password == "" {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	user.Password = authService.HashPass(user.Password)
	_, error := h.repo.CreateUser(user)

	if error != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
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
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	if err := userService.ValidateUser(userReq); err != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	userDB, err := h.repo.GetUserByEmail(userReq.Email)

	if err != nil {
		utils.ErrRes(w, errors.User404Err, http.StatusNotFound)
		return
	}

	if authService.IsPassEqual(userReq.Password, userDB.Password) {
		utils.ErrRes(w, errors.AuthIvalidPass, http.StatusBadRequest)
		return
	}

	token := authService.GenerateJwt(userReq.Email)
	tokens = append(tokens, token)

	w.WriteHeader(http.StatusOK)
	payload := model.Authenticated{Token: token.Token}
	encoder.Encode(payload)
}
