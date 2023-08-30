package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secret := []byte("som3th1ng_s3cr3t") //.ENV
			return secret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if authHeader != "" {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) == 2 {
			return splitToken[1]

		}
	}

	// If the token is not found in the Authorization header, try to get it from the "token" query parameter
	return r.URL.Query().Get("token")
}
