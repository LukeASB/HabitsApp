package middleware

import (
	"dohabits/helper"
	"dohabits/logger"
	"net/http"
)

func JSONMiddleware(logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	functionName := helper.GetFunctionName()
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog(functionName, "")
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
	}
}
