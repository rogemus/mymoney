package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	"tracker/pkg/utils"
	errors "tracker/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Register"))
}

func (h *UserHandler) Protected(w http.ResponseWriter, r *model.ProtectedRequest) {
  fmt.Printf("%v >>>> \n", r.UserEmail)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Protected"))
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		utils.ErrRes(w, errors.Generic400Err, http.StatusBadRequest)
		return
	}

	token := generateJwt(user.Email)
	tokens = append(tokens, token)

	w.WriteHeader(http.StatusOK)
	payload := model.Authenticated{Token: token.Token}
	encoder.Encode(payload)
}

func generateJwt(userEmail string) model.Token {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &model.Claims{
		UserEmail: "test@test.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenStr, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		fmt.Printf("%v", err)
	}

	token := model.Token{
		Token:     jwtTokenStr,
		UserEmail: userEmail,
	}
	return token
}
