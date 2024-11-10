package session

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokens struct {
	jwtKey []byte
}

type IJWTTokens interface {
	GetJWTKey() []byte
	GenerateAccessJWT(username string) (string, error)
	GenerateRefreshJWT(username string) (string, error)
	RefreshJWTTokens() (string, error)
	GetJWTToken(r *http.Request) (string, error)
	generateToken(username string, expirationTime time.Time) (string, error)
}

func NewJWTTokens(jwtSecret string) *JWTTokens {
	return &JWTTokens{
		jwtKey: []byte(jwtSecret),
	}
}

var refreshTokenFile = "refresh_token.txt" // Stored in txt file for now

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (sa *JWTTokens) GetJWTKey() []byte {
	return sa.jwtKey
}

// GenerateAccessJWT generates a new JWT for a given username
func (sa *JWTTokens) GenerateAccessJWT(username string) (string, error) {
	shortLivedJWT := time.Now().Add(5 * time.Minute)
	return sa.generateToken(username, shortLivedJWT)
}

// GenerateRefreshJWT generates a new long-lived refresh token
func (sa *JWTTokens) GenerateRefreshJWT(username string) (string, error) {
	longLivedJWT := time.Now().Add(24 * time.Hour) // Long-lived refresh token
	tokenString, err := sa.generateToken(username, longLivedJWT)

	// Store the refresh token in a file - temp
	if err == nil {
		err = os.WriteFile(refreshTokenFile, []byte(tokenString), 0644)
	}

	return tokenString, err
}

func (sa *JWTTokens) RefreshJWTTokens() (string, error) {
	refreshToken, err := os.ReadFile(refreshTokenFile)

	if err != nil {
		return "", err
	}

	// Check if Refresh Token is valid move to func.
	token, err := jwt.Parse(string(refreshToken), func(token *jwt.Token) (interface{}, error) {
		return sa.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	// Generate a new access token and refresh token
	newJWTAccessToken, err := sa.GenerateAccessJWT("testuser")
	if err != nil {
		return "", err
	}

	_, err = sa.GenerateRefreshJWT("testuser")
	if err != nil {
		return "", err
	}

	return newJWTAccessToken, nil
}

func (sa *JWTTokens) GetJWTToken(r *http.Request) (string, error) {
	claims, ok := r.Context().Value("claims").(*Claims)

	if !ok {
		return "", errors.New("JWTTokens.GetJWTToken - No claims value")
	}

	return claims.Username, nil
}

func (sa *JWTTokens) generateToken(username string, expirationTime time.Time) (string, error) {
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
