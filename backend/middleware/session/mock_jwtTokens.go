package session

import (
	"time"
)

var _ IJSONWebToken = (*MockJWTTokens)(nil)

type MockJWTTokens struct {
	jwtKey []byte
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
func (sa *MockJWTTokens) generateShortLivedJSONWebToken(username string) (string, error) {
	return "", nil
}

// GenerateRefreshJWT generates a new long-lived refresh token
func (sa *MockJWTTokens) generateLongLivedJSONWebToken(username string) (string, error) {
	return "", nil
}

func (sa *MockJWTTokens) HandleLongLivedJSONWebToken(username string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es", nil
}

func (sa *MockJWTTokens) generateJSONWebToken(username string, expirationTime time.Time) (string, error) {
	return "", nil
}

func (sa *MockJWTTokens) DestroyJWTRefreshToken(username string) error {
	return nil
}
