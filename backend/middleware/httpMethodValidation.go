package middleware

import (
	"dohabits/helper"
	"dohabits/logger"
	"net/http"
)

func HTTPMethodValidation(httpMethod string, logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	functionName := helper.GetFunctionName()
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog(functionName, "")

			if r.Method != httpMethod {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}
