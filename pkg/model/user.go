package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Uuid     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

type Authenticated struct {
	Token string `json:"token"`
}
