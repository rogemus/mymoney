package service

import (
	"os"
	"time"
	"tracker/pkg/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func GenerateJwt(userID int, userEmail string) model.Token {
	expirationTime := time.Now().Add(expirationDuration)
	expiresAt := jwt.NewNumericDate(expirationTime)
	claims := &model.Claims{
		UserEmail: userEmail,
		UserID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenStr, _ := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	token := model.Token{
		Token:     jwtTokenStr,
		UserEmail: userEmail,
	}
	return token
}
