package middleware

import (
	"dohabits/logger"
	"fmt"
	"net/http"
)

func ErrorHandlingMiddleware(logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog("middleware.ErrorHandlingMiddleware")

			defer func() {
				if err := recover(); err != nil {
					logger.ErrorLog(fmt.Sprintf("middleware.ErrorHandlingMiddleware - Err: %s", err))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		}
	}
}
