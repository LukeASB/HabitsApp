package session

import (
	"context"
	"dohabits/logger"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const claimsKey = contextKey("claims")

func AuthMiddleware(jwtTokens IJWTTokens, logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog("session.AuthMiddleware - Start")

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				logger.ErrorLog("session.AuthMiddleware - No Auth Header present")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtTokens.GetJWTKey(), nil
			})

			if err != nil || !token.Valid {
				logger.ErrorLog("session.AuthMiddleware - JWT Token error")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			nwAccessToken, err := jwtTokens.RefreshJWTTokens(claims.Username)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", nwAccessToken))
			ctx := context.WithValue(r.Context(), claimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
