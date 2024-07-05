package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
