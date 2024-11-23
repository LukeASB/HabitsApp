package session

import (
	"dohabits/logger"
	"net/http"
)

type IMockCSRFToken interface {
	CSRFToken(w http.ResponseWriter) error
	generateCSRFToken() (string, error)
	setCSRFToken(w http.ResponseWriter, csrfToken string)
	ValidateCSRFToken(r *http.Request) error
	DestroyCSRFToken(w http.ResponseWriter)
}
type MockCSRFToken struct {
	logger logger.ILogger
}

func NewMockCSRFToken(logger logger.ILogger) *MockCSRFToken {
	return &MockCSRFToken{logger: logger}
}

func (mw *MockCSRFToken) CSRFToken(w http.ResponseWriter) error {

	return nil
}

// GenerateCSRFToken generates a secure CSRF token
func (mw *MockCSRFToken) generateCSRFToken() (string, error) {
	return "", nil
}

// SetCSRFToken sets the CSRF token in both the response cookie and header.
func (mw *MockCSRFToken) setCSRFToken(w http.ResponseWriter, csrfToken string) {
}

// ValidateCSRFToken validates the CSRF token in the request
func (mw *MockCSRFToken) ValidateCSRFToken(r *http.Request) error {
	return nil
}

func (mw *MockCSRFToken) DestroyCSRFToken(w http.ResponseWriter) {
}
