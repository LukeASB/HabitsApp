package session

import (
	"context"
	"dohabits/helper"
	"dohabits/logger"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ClaimsKey = contextKey("claims")

func AuthMiddleware(jwtTokens IJSONWebToken, logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	functionName := helper.GetFunctionName()
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog(functionName, "")

			authHeader := r.Header.Get("Authorization")

			// Attempt to get the Access Token from the Authorization Header
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				logger.ErrorLog(functionName, "No Auth Header present")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtTokens.GetJWTKey(), nil
			})

			if err != nil || !token.Valid {
				logger.ErrorLog(functionName, "JWT Token error")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// Check Access Token Expiry Time - If there's less than a minute left get the Refresh Token to attempt to re-issue a new Access Token
			expirationTime := claims.ExpiresAt.Time

			timeRemaining := time.Until(expirationTime)

			if timeRemaining > time.Minute {
				w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
				ctx := context.WithValue(r.Context(), ClaimsKey, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Attempt to get a new Access Token
			newAccessToken, err := jwtTokens.HandleLongLivedJSONWebToken(claims.Username)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", newAccessToken))
			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
