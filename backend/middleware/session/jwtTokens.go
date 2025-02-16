package session

import (
	"dohabits/db"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokens struct {
	jwtKey []byte
	db     db.IDB
}

type IJWTTokens interface {
	GetJWTKey() []byte
	GenerateJSONWebTokens(username string) (string, string, error)
	generateAccessJWT(username string) (string, error)
	generateRefreshJWT(username string) (string, error)
	RefreshJWTTokens(username string) (string, error)
	GetJWTToken(r *http.Request) (string, error)
	generateToken(username string, expirationTime time.Time) (string, error)
}

func NewJWTTokens(jwtSecret string, db db.IDB) *JWTTokens {
	return &JWTTokens{
		jwtKey: []byte(jwtSecret),
		db:     db,
	}
}

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (sa *JWTTokens) GetJWTKey() []byte {
	return sa.jwtKey
}

// GenerateJSONWebTokens generates the Access/Refresh JWT Tokens
func (sa *JWTTokens) GenerateJSONWebTokens(username string) (string, string, error) {
	accessToken, err := sa.generateAccessJWT(username) // username will be passed as query param

	if err != nil {
		return "", "", err
	}

	refreshToken, err := sa.generateRefreshJWT(username)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
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

	return tokenString, err
}

func (sa *JWTTokens) RefreshJWTTokens(username string) (string, error) {
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
