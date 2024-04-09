package service

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
	"tracker/pkg/model"

	"github.com/golang-jwt/jwt/v5"
)

func IsPassEqual(pass_1, hashPass_2 string) bool {
	hashPass_1 := HashPass(pass_1)

	return hashPass_1 == hashPass_2
}

func HashPass(pass string) string {
	hashPass := sha256.Sum256([]byte(pass))
	return fmt.Sprintf("%x", hashPass)
}

func GenerateJwt(userEmail string) model.Token {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		UserEmail: userEmail,
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
