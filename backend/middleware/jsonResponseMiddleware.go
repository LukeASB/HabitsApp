package middleware

import (
	"dohabits/logger"
	"net/http"
)

func JSONMiddleware(logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog("middleware.JSONMiddlware")
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
	}
}
