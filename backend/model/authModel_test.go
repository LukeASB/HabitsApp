package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/middleware/session"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	testCases := []struct {
		name                string
		want                bool
		userRegisterRequest *data.RegisterUserRequest
	}{
		{
			name: "Successfully create a user",
			want: true,
			userRegisterRequest: &data.RegisterUserRequest{
				EmailAddress: "test@test.com",
				Password:     "1secret?Password",
				FirstName:    "First",
				LastName:     "Last",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			registeredUserData, err := authModel.RegisterUserHandler(val.userRegisterRequest)

			if err != nil {
				t.Errorf("TestRegisterUser Failed - err=%s", err)
				return
			}

			got := registeredUserData.Success

			if got != val.want {
				t.Errorf("TestRegisterUser Failed - got=%v, want=%v", got, val.want)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name     string
		want     bool
		userAuth *data.UserAuth
	}{
		{
			name:     "Successfully login",
			want:     true,
			userAuth: &data.UserAuth{EmailAddress: "johndoe1@example.com", Password: "1secret?Password"},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			userLoggedIn, err := authModel.LoginHandler(w, val.userAuth, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("TestLoginHandler Failed - err=%s", err)
				return
			}

			got := userLoggedIn.Success

			if got != val.want {
				t.Errorf("TestLoginHandler Failed - got=%v, want=%v", got, val.want)
			}
		})
	}
}

func TestLogoutHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name                 string
		userLoggedOutRequest *data.UserLoggedOutRequest
		want                 error
	}{
		{
			name:                 "Successfully logout",
			userLoggedOutRequest: &data.UserLoggedOutRequest{EmailAddress: "john.loggedin@example.com"},
			want:                 nil,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			var refreshTokenPath = "data/mock_refresh_tokens"
			var refreshTokenFile = "mock_refresh_token.txt"
			err := os.WriteFile(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, val.userLoggedOutRequest.EmailAddress, refreshTokenFile), []byte("testjwt"), 0644)
			if err != nil {
				t.Errorf("TestLogoutHandler Failed - err=%s", err)
			}

			err = authModel.LogoutHandler(w, val.userLoggedOutRequest, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("TestLogoutHandler Failed - err=%s", err)
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name               string
		userRefreshRequest *data.UserRefreshRequest
		want               string
	}{
		{
			name:               "Get Refresh Token",
			userRefreshRequest: &data.UserRefreshRequest{EmailAddress: "johndoe1@example.com"},
			want:               "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es",
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			newAccessToken, _, err := authModel.RefreshHandler(w, val.userRefreshRequest, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("TestRefreshHandler Failed - err=%s", err)
				return
			}

			got := newAccessToken

			if val.want != got {
				t.Errorf("TestLogoutHandler Failed - got=%s, want=%s", got, val.want)
			}
		})
	}
}
