package session

import (
	"crypto/rand"
	"dohabits/logger"
	"encoding/base64"
	"fmt"
	"net/http"
)

type ICSRFToken interface {
	GenerateCSRFToken() (string, error)
	ValidateCSRFToken(r *http.Request) error
}

type CSRFToken struct {
	logger logger.ILogger
}

func NewCSRFToken(logger logger.ILogger) *CSRFToken {
	return &CSRFToken{logger: logger}
}

// GenerateCSRFToken generates a secure CSRF token
func (mw *CSRFToken) GenerateCSRFToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		mw.logger.ErrorLog(fmt.Sprintf("Failed to generate CSRF token: %s", err.Error()))
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(tokenBytes), nil
}

// ValidateCSRFToken validates the CSRF token in the request
func (mw *CSRFToken) ValidateCSRFToken(r *http.Request) error {
	tokenFromHeader := r.Header.Get("X-CSRF-Token")
	tokenFromCookie, err := r.Cookie("csrf_token")

	if err != nil || tokenFromHeader != tokenFromCookie.Value {
		mw.logger.ErrorLog("Invalid or missing CSRF token")
		return http.ErrNoCookie // Consider using a custom error type for more clarity
	}
	return nil
}
