package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
