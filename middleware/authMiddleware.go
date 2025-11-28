package middleware

import (
	"net/http"
	"simple/config"
	"simple/utils"
	"strings"
)

var jwtAccessTokenSecret = []byte(config.LoadConfig().AccessSecretkey)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authorization header read
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header misssing", http.StatusUnauthorized)
			return
		}

		// format check: "Bearer <token>"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		// extract token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// verify and validate JWT
		_, err := utils.VerifyToken(tokenStr, jwtAccessTokenSecret)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Token valid --> continue request
		next.ServeHTTP(w, r)
	})
}
