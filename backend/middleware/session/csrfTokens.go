package session

import (
	"crypto/rand"
	"dohabits/helper"
	"dohabits/logger"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

const (
	csrf_token_cookie = "csrf_token"
	csrf_header       = "X-CSRF-Token"
)

type ICSRFToken interface {
	CSRFToken(w http.ResponseWriter) error
	generateCSRFToken() (string, error)
	setCSRFToken(w http.ResponseWriter, csrfToken string)
	ValidateCSRFToken(r *http.Request) error
	DestroyCSRFToken(w http.ResponseWriter)
}
type CSRFToken struct {
	logger logger.ILogger
}

func NewCSRFToken(logger logger.ILogger) *CSRFToken {
	return &CSRFToken{logger: logger}
}

func (mw *CSRFToken) CSRFToken(w http.ResponseWriter) error {
	csrfToken, err := mw.generateCSRFToken()

	if err != nil {
		mw.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("CSRF token generation failed: %v", err))
		return err
	}

	mw.setCSRFToken(w, csrfToken)

	return nil
}

// GenerateCSRFToken generates a secure CSRF token
func (mw *CSRFToken) generateCSRFToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		mw.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to generate CSRF token: %s", err.Error()))
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(tokenBytes), nil
}

// SetCSRFToken sets the CSRF token in both the response cookie and header.
func (mw *CSRFToken) setCSRFToken(w http.ResponseWriter, csrfToken string) {
	// Set the CSRF token as a cookie in the response
	http.SetCookie(w, &http.Cookie{
		Name:     csrf_token_cookie,
		Value:    csrfToken,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Use Secure in production if using HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	// Optionally, also include the token in the header
	w.Header().Set(csrf_header, csrfToken)
}

// ValidateCSRFToken validates the CSRF token in the request
func (mw *CSRFToken) ValidateCSRFToken(r *http.Request) error {
	tokenFromHeader := r.Header.Get(csrf_header)
	tokenFromCookie, err := r.Cookie(csrf_token_cookie)

	if err != nil || tokenFromHeader != tokenFromCookie.Value {
		mw.logger.ErrorLog(helper.GetFunctionName(), "Invalid or missing CSRF token")
		return http.ErrNoCookie // Consider using a custom error type for more clarity
	}
	return nil
}

func (mw *CSRFToken) DestroyCSRFToken(w http.ResponseWriter) {
	// Invalidate the CSRF token by clearing the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     csrf_token_cookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Use Secure in production if using HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0), // Expire the cookie immediately
	})
}
