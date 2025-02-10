package middleware

import (
	"dohabits/helper"
	"dohabits/logger"
	"fmt"
	"net/http"
)

func ErrorHandlingMiddleware(logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	functionName := helper.GetFunctionName()
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog(functionName, "")

			defer func() {
				if err := recover(); err != nil {
					logger.ErrorLog(functionName, fmt.Sprintf("middleware.ErrorHandlingMiddleware - Err: %s", err))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		}
	}
}
