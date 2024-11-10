package middleware

import (
	"dohabits/logger"
	"dohabits/middleware/session"
	"fmt"
	"net/http"
)

func CSRFToken(csrf session.ICSRFToken, logger logger.ILogger) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// If CSRF token is required for certain HTTP methods (e.g., POST, PUT, DELETE), validate it
			if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
				// Validate CSRF Token
				if err := csrf.ValidateCSRFToken(r); err != nil {
					logger.ErrorLog(fmt.Sprintf("CSRF validation failed: %v", err))
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
				}
			}

			// Generate CSRF token for the response (for use in subsequent requests)
			csrfToken, err := csrf.GenerateCSRFToken()
			if err != nil {
				logger.ErrorLog(fmt.Sprintf("CSRF token generation failed: %v", err))
				http.Error(w, "Error generating CSRF token", http.StatusInternalServerError)
				return
			}

			// Set CSRF token as HttpOnly cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "csrf_token",
				Value:    csrfToken,
				HttpOnly: true,
				// Secure:   true, // If your site uses HTTPS
				SameSite: http.SameSiteStrictMode,
			})

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		}
	}
}
