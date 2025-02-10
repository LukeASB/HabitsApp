package session

import (
	"dohabits/logger"
	"net/http"
)

var _ ICSRFToken = (*MockCSRFToken)(nil)

type MockCSRFToken struct {
	logger logger.ILogger
}

func NewMockCSRFToken(logger logger.ILogger) *MockCSRFToken {
	return &MockCSRFToken{logger: logger}
}

func (mw *MockCSRFToken) CSRFToken(w http.ResponseWriter) (string, error) {
	return "", nil
}

// GenerateCSRFToken generates a secure CSRF token
func (mw *MockCSRFToken) generateCSRFToken() (string, error) {
	return "", nil
}

// ValidateCSRFToken validates the CSRF token in the request
func (mw *MockCSRFToken) ValidateCSRFToken(r *http.Request) error {
	return nil
}
