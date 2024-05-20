package model

import (
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SessionDuration = 24 * time.Hour

var (
	Sessions   = make(map[string]Session)
	SessionMux = &sync.Mutex{}
)

type Session struct {
	Id            string
	Created       time.Time
	UserEmail     string
	ExpiresAt     time.Time
	Duration      int
}

type ProtectedRequest struct {
	*http.Request
	UserEmail string
	UserID    int
}

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
