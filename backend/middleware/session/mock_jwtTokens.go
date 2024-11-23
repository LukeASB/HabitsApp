package session

import (
	"net/http"
	"time"
)

type MockJWTTokens struct {
	jwtKey []byte
}

type IMockJWTTokens interface {
	GetJWTKey() []byte
	GenerateJSONWebTokens(username string) (string, string, error)
	generateAccessJWT(username string) (string, error)
	generateRefreshJWT(username string) (string, error)
	RefreshJWTTokens(username string) (string, error)
	GetJWTToken(r *http.Request) (string, error)
	generateToken(username string, expirationTime time.Time) (string, error)
	DestroyJWTRefreshToken(username string) error
}

func NewMockJWTTokens(jwtSecret string) *MockJWTTokens {
	return &MockJWTTokens{
		jwtKey: []byte(jwtSecret),
	}
}

func (sa *MockJWTTokens) GetJWTKey() []byte {
	return sa.jwtKey
}

// GenerateJSONWebTokens generates the Access/Refresh JWT Tokens
func (sa *MockJWTTokens) GenerateJSONWebTokens(username string) (string, string, error) {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMTczNjY2fQ.VqESTQfYNylflZ7treHfUEg8dHKo5xxwQKuyMLr6u0A"
	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NzY5fQ.avapNayXoWcdw691FrutufJycVS0FvCvOjAtmvVTk-o"
	return accessToken, refreshToken, nil
}

// GenerateAccessJWT generates a new JWT for a given username
func (sa *MockJWTTokens) generateAccessJWT(username string) (string, error) {
	return "", nil
}

// GenerateRefreshJWT generates a new long-lived refresh token
func (sa *MockJWTTokens) generateRefreshJWT(username string) (string, error) {
	return "", nil
}

func (sa *MockJWTTokens) RefreshJWTTokens(username string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es", nil
}

func (sa *MockJWTTokens) GetJWTToken(r *http.Request) (string, error) {
	return "", nil
}

func (sa *MockJWTTokens) generateToken(username string, expirationTime time.Time) (string, error) {
	return "", nil
}

func (sa *MockJWTTokens) DestroyJWTRefreshToken(username string) error {
	return nil
}
