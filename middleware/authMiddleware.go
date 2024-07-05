package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := cookie.Value
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add the username to the context for access in the handlers
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
