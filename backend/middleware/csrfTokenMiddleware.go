package middleware

import (
	"dohabits/logger"
	"dohabits/middleware/session"
	"fmt"
	"net/http"
)

func CSRFToken(csrfTokens session.ICSRFToken, logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.InfoLog("middleware.CSRFToken")
			// If CSRF token is required for certain HTTP methods (e.g., POST, PUT, DELETE), validate it
			if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
				// Validate CSRF Token
				if err := csrfTokens.ValidateCSRFToken(r); err != nil {
					logger.ErrorLog(fmt.Sprintf("CSRF validation failed: %v", err))
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
				}
			}

			if err := csrfTokens.CSRFToken(w); err != nil {
				http.Error(w, "Error generating CSRF token", http.StatusInternalServerError)
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		}
	}
}
