package service

import (
	"crypto/rand"
	"encoding/base32"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var expirationDuration = 60 * time.Minute

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSessionId() string {
	bytes := make([]byte, 15)
	rand.Read(bytes)
	sessionId := base32.StdEncoding.EncodeToString(bytes)

	return sessionId
}
