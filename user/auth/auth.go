package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   uint   `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string, id uint) (tokenString string, err error) {
	var jwtKey = []byte(os.Getenv("TOKEN_SECRET"))
	tokenTimeout, err := strconv.Atoi(os.Getenv("TOKEN_TIMEOUT"))
	duration := time.Duration(tokenTimeout) * time.Hour
	expirationTime := time.Now().Add(time.Duration(duration))
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}
