package session

import (
	"errors"
	"fmt"
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
	GenerateJSONWebTokens(username string) (string, error)
	generateAccessJWT(username string) (string, error)
	generateRefreshJWT(username string) (string, error)
	RefreshJWTTokens(username string) (string, error)
	GetJWTToken(r *http.Request) (string, error)
	generateToken(username string, expirationTime time.Time) (string, error)
	DestroyJWTRefreshToken(username string) error
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

// GenerateJSONWebTokens generates the Access/Refresh JWT Tokens
func (sa *JWTTokens) GenerateJSONWebTokens(username string) (string, error) {
	accessToken, err := sa.generateAccessJWT(username) // username will be passed as query param

	if err != nil {
		return "", err
	}

	_, err = sa.generateRefreshJWT(username)

	if err != nil {
		return "", err
	}

	return accessToken, err
}

// GenerateAccessJWT generates a new JWT for a given username
func (sa *JWTTokens) generateAccessJWT(username string) (string, error) {
	shortLivedJWT := time.Now().Add(5 * time.Minute)
	return sa.generateToken(username, shortLivedJWT)
}

// GenerateRefreshJWT generates a new long-lived refresh token
func (sa *JWTTokens) generateRefreshJWT(username string) (string, error) {
	longLivedJWT := time.Now().Add(24 * time.Hour) // Long-lived refresh token
	tokenString, err := sa.generateToken(username, longLivedJWT)

	// Store the refresh token in a file - temp
	if err == nil {
		err = os.WriteFile(fmt.Sprintf("%s_%s", username, refreshTokenFile), []byte(tokenString), 0644)
	}

	return tokenString, err
}

func (sa *JWTTokens) RefreshJWTTokens(username string) (string, error) {
	refreshToken, err := os.ReadFile(fmt.Sprintf("%s_%s", username, refreshTokenFile))

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
	newJWTAccessToken, err := sa.GenerateJSONWebTokens(username)

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

func (sa *JWTTokens) DestroyJWTRefreshToken(username string) error {
	// Stored in file for now
	if err := os.Remove(refreshTokenFile); err != nil {
		return err
	}
	return nil
}
