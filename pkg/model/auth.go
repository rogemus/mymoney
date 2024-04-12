package model

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	ID        int
	Uuid      string
	Token     string
	UserEmail string
	Created   time.Time
}

type Claims struct {
	UserEmail string
	UserID    int
	jwt.RegisteredClaims
}

type ProtectedRequest struct {
	*http.Request
	UserEmail string
	UserID    int
}
