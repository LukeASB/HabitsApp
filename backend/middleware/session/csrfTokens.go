package session

import (
	"dohabits/helper"
	"dohabits/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ICSRFToken interface {
	CSRFToken(w http.ResponseWriter) (string, error)
	generateCSRFToken() (string, error)
	ValidateCSRFToken(r *http.Request) error
}
type CSRFToken struct {
	jwtKey []byte
	logger logger.ILogger
}

func NewCSRFToken(jwtSecret string, logger logger.ILogger) *CSRFToken {
	return &CSRFToken{jwtKey: []byte(jwtSecret), logger: logger}
}

func (cs *CSRFToken) CSRFToken(w http.ResponseWriter) (string, error) {
	csrfToken, err := cs.generateCSRFToken()

	if err != nil {
		cs.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("CSRF token generation failed: %v", err))
		return "", err
	}

	return csrfToken, nil
}

// GenerateCSRFToken generates a secure CSRF token
func (cs *CSRFToken) generateCSRFToken() (string, error) {
	shortLivedJWT := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(shortLivedJWT),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(cs.jwtKey)
}

// ValidateCSRFToken validates the CSRF token in the request
func (cs *CSRFToken) ValidateCSRFToken(r *http.Request) error {
	tokenFromHeader := r.Header.Get("X-CSRF-Token")

	if tokenFromHeader == "" {
		cs.logger.ErrorLog(helper.GetFunctionName(), "Invalid or missing CSRF token")
		return fmt.Errorf("%s - Invalid or missing CSRF token", helper.GetFunctionName())
	}

	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			cs.logger.ErrorLog(helper.GetFunctionName(), "Unexpected Signing Method")
			return nil, fmt.Errorf("%s - Unexpected Signing Method", helper.GetFunctionName())
		}
		return cs.jwtKey, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Optional: Check expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				cs.logger.ErrorLog(helper.GetFunctionName(), "Invalid CSRF Token")
				return fmt.Errorf("%s - Invalid CSRF Token", helper.GetFunctionName())
			}
		}
		return nil
	}

	cs.logger.ErrorLog(helper.GetFunctionName(), "Invalid CSRF Token")
	return fmt.Errorf("%s - Invalid CSRF Token", helper.GetFunctionName())
}
