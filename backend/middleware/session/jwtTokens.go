package session

import (
	"dohabits/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JSONWebToken struct {
	jwtKey []byte
	db     db.IDB
}

type IJSONWebToken interface {
	GetJWTKey() []byte
	GenerateJSONWebTokens(username string) (string, string, error)
	generateShortLivedJSONWebToken(username string) (string, error)
	generateLongLivedJSONWebToken(username string) (string, error)
	generateJSONWebToken(username string, expirationTime time.Time) (string, error)
	HandleLongLivedJSONWebToken(username string) (string, error)
}

func NewJSONWebToken(jwtSecret string, db db.IDB) *JSONWebToken {
	return &JSONWebToken{
		jwtKey: []byte(jwtSecret),
		db:     db,
	}
}

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (sa *JSONWebToken) GetJWTKey() []byte {
	return sa.jwtKey
}

// Generates the Access/Refresh JWT Tokens
func (sa *JSONWebToken) GenerateJSONWebTokens(username string) (string, string, error) {
	accessToken, err := sa.generateShortLivedJSONWebToken(username) // username will be passed as query param

	if err != nil {
		return "", "", err
	}

	refreshToken, err := sa.generateLongLivedJSONWebToken(username)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// Generates a new short-lived access token
func (sa *JSONWebToken) generateShortLivedJSONWebToken(username string) (string, error) {
	shortLivedJWT := time.Now().Add(5 * time.Minute)
	return sa.generateJSONWebToken(username, shortLivedJWT)
}

// Generates a new long-lived refresh token
func (sa *JSONWebToken) generateLongLivedJSONWebToken(username string) (string, error) {
	longLivedJWT := time.Now().Add(24 * time.Hour) // Long-lived refresh token
	tokenString, err := sa.generateJSONWebToken(username, longLivedJWT)

	return tokenString, err
}

// Generate a JSON Web Token with username and expiry claims
func (sa *JSONWebToken) generateJSONWebToken(username string, expirationTime time.Time) (string, error) {
	expires := expirationTime

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(sa.jwtKey)
}

// Retrieve the refresh token and validate it
func (sa *JSONWebToken) HandleLongLivedJSONWebToken(username string) (string, error) {
	refreshToken, err := sa.db.RetrieveUserSession(username, "")

	if err != nil {
		return "", err
	}

	// Check if Refresh Token is valid
	token, err := jwt.Parse(string(refreshToken), func(token *jwt.Token) (interface{}, error) {
		return sa.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	// Generate a new access token and refresh token
	newJWTAccessToken, _, err := sa.GenerateJSONWebTokens(username)

	if err != nil {
		return "", err
	}

	return newJWTAccessToken, nil
}
