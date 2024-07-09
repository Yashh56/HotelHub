package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/Yashh56/HotelHub/models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Debug statements
		log.Printf("Parsed token: %v", token)
		log.Printf("UserID from token: %v", claims.UserID)

		ctx := context.WithValue(r.Context(), "userId", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
