package model

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Token     string
	UserEmail string
}

type Tokens []Token

type Claims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}

type ProtectedRequest struct {
	*http.Request
	UserEmail string
}
